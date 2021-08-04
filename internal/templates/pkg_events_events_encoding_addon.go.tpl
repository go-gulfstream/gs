package addon

func init() {
       {{ if .Event.Payload -}}
           gulfstreamevent.RegisterCodec({{.OutEvent.Name}}, &{{.OutEvent.Payload}}{})
       {{end -}}
}

{{ if .OutEvent.Payload -}}
            func (c *{{.OutEvent.Payload}}) MarshalBinary() ([]byte, error) {
            	return json.Marshal(c)
            }

            func (c *{{.OutEvent.Payload}}) UnmarshalBinary(data []byte) error {
            	return json.Unmarshal(data, c)
            }
{{end}}