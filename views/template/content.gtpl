{{define "content"}}
    {{template "header"}}
    <h1>ネストのデモ</h1>
    <ul>
        <li>ネストではdefineを使用してサブテンプレートを定義します。</li>
        <li>templateの使用をコール</li>
    </ul>
    {{template "footer"}}
{{end}}
