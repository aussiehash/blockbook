{{define "main" -}}
ARCHIVE := $(shell basename {{.Backend.BinaryURL}})

all:
	wget {{.Backend.BinaryURL}}
{{- if eq .Backend.VerificationType "gpg"}}
	wget {{.Backend.VerificationSource}} -O checksum
	gpg --verify checksum
{{- else if eq .Backend.VerificationType "gpg-sha256"}}
	wget {{.Backend.VerificationSource}} -O checksum
	gpg --verify checksum
	sha256sum -c --ignore-missing checksum
{{- else if eq .Backend.VerificationType "sha256"}}
	[ "$$(sha256sum backend.tar.gz | cut -d ' ' -f 1)" = "{{.Backend.VerificationSource}}" ]
{{- end}}
	mkdir backend
	tar -xf ${ARCHIVE} -C backend --strip 1
{{- range $index, $name := .Backend.ExcludeFiles}}
	rm backend/{{$name}}
{{- end}}

clean:
	rm -rf backend
	rm -f ${ARCHIVE}
	rm -f checksum
{{end}}
