package lib

const (
	E = iota
	W
	I
)

const (
	NAN              = "%s n'est pas un nombre"
	NEGATIVESQRT     = "Racine d'un nombre négatif"
	NEGATIVEFACT     = "Factorielle d'un nombre négatif"
	DECIMALFACT      = "Factorielle d'un nombre décimal"
	DIVIDEBYZERO     = "Division par zéro"
	UNKNOWN          = "Erreur inconnue"
	EMPTYSTACK       = "Pile vide"
	EMPTYSTACKN      = "La pile contient moins de %d entrées"
	EMPTYMEMORY      = "Mémoire vide"
	EMPTYREGISTER    = "Registre '%s' vide"
	EMPTYREGISTERN   = "Le registre '%s' contient moins de %d entrées"
	NONEMPTYREGISTER = "Le registre '%s' n'est pas vide."
	ERASEREGISTER    = "Voulez-vous écraser ce registre ? [O/n/a]"
	CANCEL           = "Action annulée..."
	CREATEDREGISTER  = "Registre '%s' alimenté..."
	UNKNOWNCOMMAND   = "Commande inconnue : %s"
	NOSECTION        = "Aucune section valide sélectionnée"
)

type Message struct {
	tpe uint64
	msg string
}

func Error(msg string) *Message {
	return &Message{E, msg}
}

func Warning(msg string) *Message {
	return &Message{W, msg}
}

func Info(msg string) *Message {
	return &Message{I, msg}
}

func (m *Message) IsError() bool {
	return m.tpe == E
}

func (m *Message) IsWarning() bool {
	return m.tpe == W
}

func (m *Message) IsInfo() bool {
	return m.tpe == I
}
