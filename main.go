package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<!doctype html>
			<html>
			<head>
				<title>Page home</title>
			</head>
			<body>
			
			<h1 align="center">Kubernetes Security Detection System</h1>
			
			<p>
			<div style="text-align:center; height:40px">
				<button onclick="resources_monitoring()">Resources Monitoring</button>
			</div>
			</p>
			<p>
			<div style="text-align:center; height:40px">
				<button onclick="property_inventory()">Property Inventory</button>
			</div>
			</p>
			<p>
			<div style="text-align:center; height:40px">
				<button onclick="logging_system()">Logging System</button>
			</div>
			</p>
			<p>
			<div style="text-align:center; height:40px">
				<button onclick="k8s_CIS()">k8s CIS Benchmark Detection</button>
			</div>
			</p>
			<p>
			<div style="text-align:center; height:40px">
				<button onclick="docker_CIS()">docker CIS Benchmark Detection</button>
			</div>
			</p>

			<p>
			<div style="text-align:center; height:40px">
				<button onclick="image_count()">image count</button>
			</div>
			</p>

			<p>
			<div style="text-align:center; height:40px">
				<button onclick="registry_count()">registry count</button>
			</div>
			</p>
			
			<script>
				function resources_monitoring() {
					window.open("http://192.168.80.240:32332");
				}
				function property_inventory() {
					window.open("http://192.168.80.240:32582");
				}
				function logging_system() {
					window.open("http://192.168.80.240:30238");
				}
				function k8s_CIS() {
					window.open("http://192.168.80.240/kube-bench-report.txt");
				}
				function docker_CIS() {
					window.open("http://192.168.80.240/docker-bench-report.txt");
				}
				function image_count() {
					window.open("http://192.168.80.240/image_count.txt");
				}
				function registry_count() {
					window.open("http://192.168.80.242/registry.txt");
				}
			</script>
			
			</body>
			</html>`
	fmt.Fprintln(w, html)
}


func main() {
	http.HandleFunc("/", home)

	err := http.ListenAndServe(":3500", nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}