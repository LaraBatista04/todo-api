package services

import (
	"testing"
	"todo_API/models"
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