// Copyright 2018 The Go Cloud Development Kit Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The deploy program builds the Guestbook server locally and deploys it to
// GKE.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gcp/deploy: ")
	myinventoryDir := flag.String("myinventory_dir", filepath.Join("..", "..", "cmd"), "directory containing the myinventory example")
	tfStatePath := flag.String("tfstate", "terraform.tfstate", "path to terraform state file")
	flag.Parse()
	if err := deploy(*myinventoryDir, *tfStatePath); err != nil {
		log.Fatal(err)
	}
}

func deploy(myinventoryDir, tfStatePath string) error {
	type tfItem struct {
		Sensitive bool
		Type      string
		Value     string
	}
	type state struct {
		Project          tfItem
		ClusterName      tfItem `json:"cluster_name"`
		ClusterZone      tfItem `json:"cluster_zone"`
		Bucket           tfItem
		DatabaseInstance tfItem `json:"database_instance"`
		DatabaseRegion   tfItem `json:"database_region"`
		MotdVarConfig    tfItem `json:"motd_var_config"`
		MotdVarName      tfItem `json:"motd_var_name"`
	}
	tfStateb, err := runb("terraform", "output", "-state", tfStatePath, "-json")
	if err != nil {
		return err
	}
	var tfState state
	if err = json.Unmarshal(tfStateb, &tfState); err != nil {
		return fmt.Errorf("parsing terraform state JSON: %v", err)
	}
	zone := tfState.ClusterZone.Value
	if zone == "" {
		return fmt.Errorf("empty or missing cluster_zone in %s", tfStatePath)
	}
	tempDir, err := ioutil.TempDir("", "myinventory-k8s-")
	if err != nil {
		return fmt.Errorf("making temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Fill in Kubernetes template parameters.
	proj := strings.Replace(tfState.Project.Value, ":", "/", -1)
	imageName := fmt.Sprintf("gcr.io/%s/myinventory", proj)
	gbyin, err := ioutil.ReadFile(filepath.Join(".", "myinventory.yaml.in"))
	if err != nil {
		return fmt.Errorf("reading myinventory.yaml.in: %v", err)
	}
	gby := string(gbyin)
	replacements := map[string]string{
		"{{IMAGE}}":             imageName,
		"{{database_instance}}": tfState.DatabaseInstance.Value,
		"{{database_region}}":   tfState.DatabaseRegion.Value,
	}
	for old, new := range replacements {
		gby = strings.Replace(gby, old, new, -1)
	}
	if err = ioutil.WriteFile(filepath.Join(tempDir, "myinventory.yaml"), []byte(gby), 0666); err != nil {
		return fmt.Errorf("writing myinventory.yaml: %v", err)
	}

	// Build Guestbook Docker image.
	log.Printf("Building %s...", imageName)
	build := exec.Command("go", "build", "-o", filepath.Join("..", "deploy", "gcp", "myinventory"))
	env := append(build.Env, "GOOS=linux", "GOARCH=amd64")
	env = append(env, os.Environ()...)
	build.Env = env
	absDir, err := filepath.Abs(myinventoryDir)
	if err != nil {
		return fmt.Errorf("getting abs path to myinventory dir (%s): %v", myinventoryDir, err)
	}
	build.Dir = absDir
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		return fmt.Errorf("building myinventory app by running %v: %v", build.Args, err)
	}
	gcp := gcloud{projectID: tfState.Project.Value}
	cbs := gcp.cmd("builds", "submit", "-t", imageName, filepath.Join(myinventoryDir, "..", "deploy", "gcp"))
	if err := cbs.Run(); err != nil {
		return fmt.Errorf("building container image with %v: %v", cbs.Args, err)
	}

	// Run on Kubernetes.
	log.Printf("Deploying to %s...", tfState.ClusterName.Value)
	getCreds := gcp.cmd("container", "clusters", "get-credentials", "--zone", zone, tfState.ClusterName.Value)
	getCreds.Stderr = os.Stderr
	if err := getCreds.Run(); err != nil {
		return fmt.Errorf("getting credentials with %v: %v", getCreds.Args, err)
	}
	kubeCmds := [][]string{
		{"kubectl", "apply", "-f", filepath.Join(tempDir, "myinventory.yaml")},
		// Force pull the latest image.
		{"kubectl", "scale", "--replicas", "0", "deployment/myinventory"},
		{"kubectl", "scale", "--replicas", "1", "deployment/myinventory"},
	}
	for _, kcmd := range kubeCmds {
		cmd := exec.Command(kcmd[0], kcmd[1:]...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("running %v: %v", cmd.Args, err)
		}
	}

	// Wait for endpoint then print it.
	log.Printf("Waiting for load balancer...")
	for {
		outb, err := runb("kubectl", "get", "service", "myinventory", "-o", "json")
		if err != nil {
			return err
		}
		var s service
		if err := json.Unmarshal(outb, &s); err != nil {
			return fmt.Errorf("parsing JSON output: %v", err)
		}
		i := s.Status.LoadBalancer.Ingress
		if len(i) == 0 || i[0].IP == "" {
			dt := time.Second
			log.Printf("No ingress returned in %s. Trying again in %v", "", dt)
			time.Sleep(dt)
			continue
		}
		endpoint := i[0].IP
		log.Printf("Deployed at http://%s:8089", endpoint)
		break
	}
	return nil
}

type service struct{ Status *status }
type status struct{ LoadBalancer loadBalancer }
type loadBalancer struct{ Ingress []ingress }
type ingress struct{ IP string }

type gcloud struct {
	projectID string
}

func (gcp *gcloud) cmd(args ...string) *exec.Cmd {
	args = append([]string{"--quiet", "--project", gcp.projectID}, args...)
	cmd := exec.Command("gcloud", args...)
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Stderr = os.Stderr
	return cmd
}

func run(args ...string) (stdout string, err error) {
	stdoutb, err := runb(args...)
	return strings.TrimSpace(string(stdoutb)), err
}

func runb(args ...string) (stdout []byte, err error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)
	stdoutb, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("running %v: %v", cmd.Args, err)
	}
	return stdoutb, nil
}
