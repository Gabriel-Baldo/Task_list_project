package main

import (
	"fmt"
	"os"
	"strings"
)

func lobby() {
	fmt.Printf("\n1. Adicionar tarefas(para dar space, use '_')\n")
	fmt.Printf("2. Listar todas as tarefas\n")
	fmt.Printf("3. Marcar tarefas como concluidas\n")
	fmt.Printf("4. Remover tarefas\n")
	fmt.Printf("0. Sair\n")
	fmt.Printf("ESCOLHA UMA ACAO:")
}

func add(tarefa string) {
	tarefinhas, _ := os.OpenFile("Tarefolas.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Printf("\n[ ] ")
	fmt.Scanf("%s", &tarefa)
	tarefinhas.WriteString("\n[ ]" + tarefa)
	defer tarefinhas.Close()
}

func listar() {
	tarefaslistadas, _ := os.ReadFile("Tarefolas.txt")
	fmt.Printf("Tarefas:\n")
	fmt.Printf(string(tarefaslistadas) + "\n")
}

func marcarcomoconcluidas() {
	tarefasconcluidas, _ := os.ReadFile("Tarefolas.txt")
	linhas := strings.Split(string(tarefasconcluidas), "\n")

	for numdalinha, linha := range linhas {
		fmt.Printf("%d. %s\n", numdalinha, linha)
	}

	var linhamarcada int
	fmt.Print("Digite o número da tarefa concluída: ")
	fmt.Scanln(&linhamarcada)

	if strings.HasPrefix(linhas[linhamarcada], "[ ]") {
		linhas[linhamarcada] = "[x]" + linhas[linhamarcada][3:]
	}

	marcacaodatarefa := strings.Join(linhas, "\n")
	os.WriteFile("Tarefolas.txt", []byte(marcacaodatarefa), 0644)

	fmt.Println("Tarefa marcada como concluída.")
}

func remover() {
	tarefasconcluidas, _ := os.ReadFile("Tarefolas.txt")
	linhas := strings.Split(string(tarefasconcluidas), "\n")

	for numdalinha, linha := range linhas {
		fmt.Printf("%d. %s\n", numdalinha, linha)
	}

	var linhamarcada int
	fmt.Print("Digite o número da tarefa que quer que seja removida: ")
	fmt.Scanln(&linhamarcada)

	if linhamarcada >= 0 && linhamarcada < len(linhas) {
		linhas = append(linhas[:linhamarcada], linhas[linhamarcada+1:]...)
		_ = os.WriteFile("Tarefolas.txt", []byte(strings.Join(linhas, "\n")), 0644)
		fmt.Printf("Tarefa removida com sucesso.")
	}
}

func limparterminal() {
	for enter := 1; enter <= 50; enter++ {
		fmt.Printf("\n")
	}
}

func main() {
	for {
		lobby()
		var acao int
		var tarefa string

		fmt.Scanf("%d", &acao)
		limparterminal()

		switch acao {
		case 1:
			add(tarefa)
			limparterminal()
		case 2:
			listar()
		case 3:
			marcarcomoconcluidas()
			limparterminal()
		case 4:
			remover()
			limparterminal()
		case 0:	
			return
		
		}
	}
}
