package lib

import (
	u "util"
)

type Option struct {
	name string
	help string
}

type Section struct {
	title   string
	options []Option
}

var SectionHelp = Section{
	title: "Affichage de l'aide",
	options: []Option{
		Option{"h", "Affiche toute l'aide"},
		Option{"hr", "Affiche la liste des rubriques et leurs id"},
		Option{"h{ids}", "Affiche les sections correspondantes aux ids spécifiés"},
	},
}

var SectionNumber = Section{
	title: "Format des nombres",
	options: []Option{
		Option{"(+|-)\\d+", "Nombre entier"},
		Option{"(+|-)(\\d+.\\d*|.\\d+)", "Nombre Décimal"},
		Option{"<entier|decimal>(E|e)<entier>", "Nombre Décimal en notation scientifique"},
		Option{"0%[01]+", "Nombre en base 2"},
		Option{"1%[01]+", "Nombre en base 2 négatif"},
		Option{"0x[\\dA-Fa-f]+", "Nombre en base 16"},
		Option{"1x[\\dA-Fa-f]+", "Nombre en base 16 négatif"},
	},
}

var SectionPrint = Section{
	title: "Options d'impression",
	options: []Option{
		Option{"p", "Impression du dernier élément de la pile"},
		Option{"e", "Impression du dernier élément de la pile au format scientifique"},
		Option{"P", "Impression de toute la pile"},
		Option{"E", "Impression de toute la pile au format scientifique"},
		Option{":P{r}", "Impression du registre 'r' de la mémoire"},
		Option{":E{r}", "Impression du registre 'r' de la mémoire au format scientifique"},
		Option{":P", "Impression de toute la mémoire"},
		Option{":E", "Impression de toute la mémoire au format scientifique"},
	},
}

var SectionArithmetique1 = Section{
	title: "Arithmétique sur 1 élément",
	options: []Option{
		Option{"|", "Valeur absolue du dernier élément de la pile"},
		Option{"v", "Racine carrée du dernier élément de la pile"},
		Option{"!", "Factorielle du dernier élément de la pile"},
		Option{"cI", "Conversion du dernier élément de la pile en nombre entier"},
		Option{"cD", "Conversion du dernier élément de la pile en nombre décimal"},
		Option{"c2", "Conversion du dernier élément de la pile en nombre en base 2"},
		Option{"c10", "Conversion du dernier élément de la pile en nombre en base 10"},
		Option{"c16", "Conversion du dernier élément de la pile en nombre en base 16"},
	},
}

var SectionArithmetique2 = Section{
	title: "Arithmétique sur 2 éléments",
	options: []Option{
		Option{"+", "Additionne les deux derniers éléments de la pile"},
		Option{"-", "Soustrait les deux derniers éléments de la pile"},
		Option{"*", "Multiplie les deux derniers éléments de la pile"},
		Option{"/", "Divise l'avant-dernier élément par le dernier élément de la pile"},
		Option{"%", "Calcule le modulo de l'avant-dernier élément par le dernier élément de la pile"},
		Option{"~", "Calcule la quotient et le reste des deux derniers éléments de la pile"},
		Option{"^", "Élève l'avant-dernier élément à la puissance du dernier élément de la pile"},
		Option{"<=>", "Compare les deux derniers éléments de la pile : -1 si p[-2] < p[-1], 0 si p[-2] = p[-1], 1 sinon"},
		Option{"=", "Retourne 1 si p[-2] = p[-1], 0 sinon"},
		Option{"<>", "Retourne 1 si p[-2] ≠ p[-1], 0 sinon"},
		Option{">=", "Retourne 1 si p[-2] ≥ p[-1], 0 sinon"},
		Option{"<=", "Retourne 1 si p[-2] ≤ p[-1], 0 sinon"},
		Option{">", "Retourne 1 si p[-2] > p[-1], 0 sinon"},
		Option{"<", "Retourne 1 si p[-2] < p[-1], 0 sinon"},
	},
}

var SectionPile = Section{
	title: "Gestion de la pile",
	options: []Option{
		Option{"c", "Suppression du dernier élément de la pile"},
		Option{"d", "Duplication du dernier élément de la pile"},
		Option{"r", "Inversion de l'ordre des deux derniers éléments de la pile"},
		Option{"C", "Réinitialisation de toute la pile"},
		Option{"D", "Duplication de toute la pile"},
		Option{"R", "Inversion de l'ordre de tous les éléments de la pile"},
	},
}

var SectionSwitch = Section{
	title: "Transferts entre la pile et la mémoire",
	options: []Option{
		Option{":s{r}", "Supprime le dernier élément de la pile et l'ajoute dans le registre 'r'"},
		Option{":l{r}", "Ajoute le dernier élément du registre 'r' dans la mémoire"},
		Option{":S{r}", "Supprime tous les éléments de la pile et les ajoute dans le registre 'r'"},
		Option{":L{r}", "Ajoute tous les éléments du registre 'r' dqns la pile"},
		Option{":={r}", "Si l'opération '=' donne 1, fait un :L{r}"},
		Option{":<>{r}", "Si l'opération '<>' donne 1, fait un :L{r}"},
		Option{":>={r}", "Si l'opération '>=' donne 1, fait un :L{r}"},
		Option{":<={r}", "Si l'opération '<=' donne 1, fait un :L{r}"},
		Option{":>{r}", "Si l'opération '>' donne 1, fait un :L{r}"},
		Option{":<{r}", "Si l'opération '<' donne 1, fait un :L{r}"},
	},
}

var SectionMemoire = Section{
	title: "Gestion de la mémoire",
	options: []Option{
		Option{":c{r}", "Suppression du dernier élément du registre 'r'; s'il n'y a plus d'élément, le registre est supprimé"},
		Option{":d{r}", "Duplication du dernier élément du registre 'r'"},
		Option{":r{r}", "Inversion de l'ordre des deux derniers éléments du registre 'r'"},
		Option{":C{r}", "Suppression du registre r"},
		Option{":D{r}", "Duplication de tout le registre r"},
		Option{":R{r}", "Inversion de l'ordre de tous les éléments du registre r"},
	},
}

var SectionOptions = Section{
	title: "Gestion des options",
	options: []Option{
		Option{"k", "Supprime le dernier élément de la pile et s'en sert comme précision automatique"},
		Option{"K", "Supprime le dernier élément de la pile et s'en sert comme précision fixe"},
		Option{"o", "Imprime les options de précision"},
	},
}

var SectionChaine = Section{
	title: "Gestion des chaînes de caractères",
	options: []Option{
		Option{"[{str}]", "Ajoute 'str' en tant que macro, sans l'évaluer, dans la pile"},
		Option{"#", "Marque le début d'un commentaire : tout ce qui suit est ignoré"},
		Option{"a", "Convertit le dernier élément de la pile en macro"},
		Option{"x", "Exécute le dernier élément de la pile si c'est une macro, sinon, ne fait rien"},
	},
}

var Sections = []Section{
	SectionHelp,
	SectionNumber,
	SectionPrint,
	SectionArithmetique1,
	SectionArithmetique2,
	SectionPile,
	SectionSwitch,
	SectionMemoire,
	SectionOptions,
	SectionChaine,
}

func printTitle(s Section) {
	HTitlePrint(s.title)
}

func printOption(o Option) {
	HOptionPrint(o.name, o.help)
}

func printSection(s Section) {
	printTitle(s)
	for _, o := range s.options {
		printOption(o)
	}
}

func PrintRubriques() {
	for i, s := range Sections {
		HTitlePrint(u.Format("(%d) %s", i, s.title))
	}
}

func PrintSections(ids []int) (ok bool) {
	l := len(Sections)
	for _, i := range ids {
		if i >= 0 && i < l {
			ok = true
			printSection(Sections[i])
		}
	}
	return
}

func PrintAll() {
	for _, s := range Sections {
		printSection(s)
	}
}
