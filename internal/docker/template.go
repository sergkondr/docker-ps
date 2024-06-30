package docker

var containerOutput = `{{ .Name }}
    Container ID:    {{ .ID }}
    Image:           {{ .Image }}
    Command:         {{ .Command }}
    Created:         {{ .CreatedTime }} {{ if .Mounts }}
    Mounts:          {{ .Mounts }}
    Network:         {{ .Network }}     {{ end }}{{ if .IPAddresses }}
    IP-address:      {{ .IPAddresses }} {{ end }}{{ if .Ports }}
    Ports:           {{ .Ports }}       {{ end }}
    Status:          {{ .Status }}
`