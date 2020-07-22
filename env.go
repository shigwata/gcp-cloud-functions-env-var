package env

import (
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
)

func Env(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "## Env")
	for _, e := range os.Environ() {
		// fmt.Println(e)
		fmt.Fprintln(w, e)
	}

	fmt.Fprintln(w, "")

	if !metadata.OnGCE() {
		fmt.Fprintln(w, "this process is not running on GCE")
		return
	}

	fmt.Fprintln(w, "## Metadata")

	fmt.Fprintln(w, "### project")

	plist := []string{
		"project/attributes/?recursive=true",
		"project/attributes/disable-legacy-endpoints",
		"project/attributes/enable-oslogin",
		"project/attributes/vmdnssetting",
		"project/attributes/ssh-keys",
		"project/numeric-project-id",
		"project/project-id",
	}
	for _, v := range plist {
		s, _ := metadata.Get(v)
		fmt.Fprintf(w, "%s: %s\n", v, s)
	}

	fmt.Fprintln(w, "### instance")

	ilist := []string{
		"instance/attributes/?recursive=true",
		"instance/attributes/enable-oslogin",
		"instance/attributes/vmdnssetting",
		"instance/attributes/ssh-keys",
		"instance/cpu-platform",
		"instance/description",
		"instance/disks/?recursive=true",
		"instance/guest-attributes/?recursive=true",
		"instance/hostname",
		"instance/id",
		"instance/machine-type",
		"instance/name",
		"instance/network-interfaces/?recursive=true",
		"instance/network-interfaces/0/access-configs/0/external-ip",
		"instance/network-interfaces/0/forwarded-ips/?recursive=true",
		"instance/scheduling/?recursive=true",
		"instance/scheduling/on-host-maintenance",
		"instance/scheduling/automatic-restart",
		"instance/scheduling/preemptible",
		"instance/maintenance-event",
		"instance/service-accounts/?recursive=true",
		"instance/service-accounts/default/email",
		"instance/service-accounts/default/token",
		"instance/service-accounts/default/identity",
		"instance/tags",
		"instance/zone",
	}
	for _, v := range ilist {
		s, _ := metadata.Get(v)
		fmt.Fprintf(w, "%s: %s\n", v, s)
	}
}
