package card

var List = map[string]string{
	"card": `{{define "card"}}
    <div class="card">
        <div class="card-body">
            <div class="card-index">
                <div class="card-top">
                    <div class="card-meta">
                        <div class="card-title">
                            <span>{{.Title}}</span>
                            <span class="card-title-action">
                                {{.Action}}
                            </span>
                        </div>
                        <div class="card-subtitle"><span>{{.SubTitle}}</span></div>
                    </div>
                </div>
                <div class="card-content">
                    {{.Content}}
                </div>
                <div class="card-footer">
                    {{.Footer}}
                </div>
            </div>
        </div>
    </div>
{{end}}`,
}
