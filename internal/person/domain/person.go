package domain

import "time"

// Person é a identidade de alguém no sistema (paciente, médico, funcionário...).
// Os papéis (Patient, Doctor, Employee) ficam em outros módulos e apontam
// para Person.ID.
type Person struct {
	ID            string
	Name          string
	BirthDate     time.Time
	MaritalStatus MaritalStatus
	Nationality   string // código ISO 3166-1 alpha-2, ex.: "BR"
	Race          Race
	Documents     Documents
	CNS           CNS

	// Contacts guarda no máximo um complemento por tipo
	// (residencial, comercial, responsável, filiação 1 e 2).
	Contacts map[ContactType]Contact
}

// Documents agrupa os documentos da pessoa.
type Documents struct {
	CPF      CPF
	RG       *RG       // opcional
	Passport *Passport // obrigatório quando Nationality != "BR"
}

// CPF e CNS são tipos próprios para receber validação depois.
type (
	CPF string
	CNS string
)

// RG agrupa os dados do registro geral.
type RG struct {
	Number      string
	Issuer      string // órgão emissor, ex.: "SSP"
	IssuerState string // UF emissora, ex.: "PA"
	IssueDate   time.Time
}

// Passport é exigido de pessoas com nacionalidade diferente de brasileira.
type Passport struct {
	Number string
}
