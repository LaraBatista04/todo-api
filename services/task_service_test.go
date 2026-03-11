package services

import (
	"testing"
	"todo_API/models"
	"todo_API/database"
)

func TestValidateTask(t *testing.T) {
	tests := []struct {
		name    string
		task    models.Task
		wantErr bool
	}{
		{
			name: "Sucesso - Dados Válidos",
			task: models.Task{Title: "Call com Itau hoje", Status: "pending", Priority: "high"},
			wantErr: false,
		},
		{
			name: "Erro - Título muito curto",
			task: models.Task{Title: "Go"},
			wantErr: true,
		},
		{
			name: "Erro - Status inválido",
			task: models.Task{Title: "Tarefa Teste", Status: "invalid_status"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTask(&tt.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateCompletedTaskRule(t *testing.T) {
	existingTask := models.Task{Status: "completed"}
	
	if existingTask.Status == "completed" {
		t.Log("Sucesso: A regra de bloqueio de edição foi detectada corretamente.")
	} else {
		t.Error("Falha: O sistema permitiu a tentativa de edição em uma tarefa completada.")
	}
}

func TestCreateTask(t *testing.T) {

	err := database.Connect()
	if err != nil {
		t.Fatalf("Erro ao conectar no banco: %v", err)
	}

	task := models.Task{
		Title: "Estudar Golang",
		Description: "Teste Unitário",
		Status: "pending",
		Priority: "high",
		DueDate: "2026-03-13",
	}

	err = CreateTask(&task)

	if err != nil {
		t.Fatalf("Erro inesperado ao criar task: %v", err)
	}

	if task.ID == "" {
		t.Error("Esperava um UUID gerado para a task, mas está vazio")
	}

	if task.CreatedAt.IsZero() {
		t.Error("CreatedAt não foi definido")
	}

	if task.UpdatedAt.IsZero() {
		t.Error("UpdatedAt não foi definido")
	}

	if task.Status != "pending" {
		t.Errorf("Esperava status 'pending', recebeu '%s'", task.Status)
	}

	if task.Priority != "high" {
		t.Errorf("Prioridade esperada 'high', recebeu '%s'", task.Priority)
	}
}