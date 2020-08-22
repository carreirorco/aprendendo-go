# Exercícios propostos no livro

Exercício 1.1: Modifique o programa **echo** para exibir também `os.Args[0]`, que é o nome do comando que o chamou.

Resolução: 8632c96

---

Exercício 1.2: Modifique o programa **echo** para exibir o `índice` e o `valor` de cada um de seus argumentos, um por linha.

Resolução: 837b1c0

---

Exercício 1.3: Experimente medir a diferença de tempo de execução entre  nossas versões potencialmente ineficientes e a versão que usa `strings.Join`. 

Resolução: 93d4b40

```
# for i in {1..3} ; do go build -o bin/echo${i} _echo${i}/echo${i}.go ; done

~/git-repos/go(master *) $ ./bin/echo1 Teste teste2 test3

Índice: 1
Valor: Teste

Índice: 2
Valor: teste2

Índice: 3
Valor: test3

 ./bin/echo1 Teste teste2 test3
Tempo: 0.00012s
~/git-repos/go(master *) $ ./bin/echo2 Teste teste2 test3
Teste teste2 test3

Tempo: 0.00006s
~/git-repos/go(master *) $ ./bin/echo3 Teste teste2 test3
Teste teste2 test3

Tempo: 0.00003s
~/git-repos/go(master *) $ 
```

---
