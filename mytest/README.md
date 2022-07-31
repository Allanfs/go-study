# MyTest

Executar cobertura de teste
---

Executar os testes e fazer output do resultado.
```
go test -coverprofile=cover.out
```

Ver cobertura em HTML
---

Abrir em HTML a cobertura de código indicando as linhas cobertas.
```
go tool covert -html=cover.out
```

VSCode e Cobertura
---

Voce pode configurar o vscode para exibir a cobertura:

`Ctrl + Shift + P` > Apply Cover Profile, e informe o caminho absoluto do arquivo de coverage.
