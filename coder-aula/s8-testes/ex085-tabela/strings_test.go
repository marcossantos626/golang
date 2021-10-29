package strings

const msgIndex = "%s (parte: %s)- Índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	testes := [] struct {
		texto string
		parte string
		esperando int
	}{
		{"Cod3r é show", "Cod3r", 0}
		{"", "", 0}
		{"Opa", "Opa", -1},
		{"Leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(mesgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}