# TDD: Ajuste da heurística e reorganização dos testes (GH-001)

## Objective & Scope

- **What:** Ajustar os pesos da heurística de detecção de domínios suspeitos e reorganizar os testes em um diretório `tests/` dedicado, com casos nomeados e dados de calibração.
- **Why:** A heurística atual marca domínios benignos como suspeitos (falsos positivos). Ao elevar o peso de agrupamentos de consoantes e tornar o bônus de caracteres únicos mais restritivo (exigindo também entropia alta), reduzimos falsos positivos sem perder a detecção de domínios DGA. Separar os testes da raiz organizationalmente facilita a manutenção e permite medir precisão/recall.
- **File Target:**Este documento está destinado a `specs/tdd-GH-001-ajusta-heuristica-e-testes.md`.

## Proposed Technical Strategy

- **Logic Flow:**
  1. Em `internal/scoring/heuristic.go`, aumentar o peso de `ConsonantClusters` de `5` para `7`.
  2. Substituir a condição `UniqueCharRatio > 0.7` por `UniqueCharRatio > 0.85 && Entropy > 3.5`, elevando o multiplicador de `50` para `60`. Assim o bônus só é aplicado quando há muitos caracteres únicos **e** alta entropia — típico de DGA, não de palavras reais.
  3. Remover `detector_test.go` da raiz e criar `tests/detector_test.go` com uma tabela de casos nomeados (domínios suspeitos e benignos).
  4. Adicionar `tests/calibrated_test.go` com limiar `WithThreshold(50)`, carregando listas de domínios benignos e maliciosos e calculando TP/FN/FP/TN, precisão e recall.
  5. Adicionar os arquivos de dados `tests/data/benign.txt` e `tests/data/malicious.txt`.
- **Impacted Files:**
  - `internal/scoring/heuristic.go` (modificado)
  - `detector_test.go` (removido da raiz)
  - `tests/detector_test.go` (novo)
  - `tests/calibrated_test.go` (novo)
  - `tests/data/benign.txt` (novo)
  - `tests/data/malicious.txt` (novo)
  - `specs/tdd-GH-001-ajusta-heuristica-e-testes.md` (este documento)
  - Nenhum diretório `doc/` (singular) será criado.
- **Guardrails (Go):**
  - Manter a assinatura `Calculate(domain string, f features.Features) float64` inalterada.
  - Usar `float64` nos cálculos para evitar erros de arredondamento.
  - Testes com `go test ./...` devem passar; tratar erros de leitura de arquivo explicitamente.

## Implementation Plan

- **Pseudocode / assinaturas:**

```go
// internal/scoring/heuristic.go
if f.ConsonantClusters >= 4 {
    score += float64(f.ConsonantClusters) * 7
}

if f.UniqueCharRatio > 0.85 && f.Entropy > 3.5 {
    score += (f.UniqueCharRatio - 0.85) * 60
}
```

```go
// tests/detector_test.go — casos nomeados
tests := []struct {
    domain        string
    expectSuspect bool
}{
    {"xkqzrt123.xyz", true},
    {"google.com", false},
    // ...
}
```

```go
// tests/calibrated_test.go — medição de precisão/recall
d := detector.New(detector.WithThreshold(50))
// carrega benign.txt / malicious.txt, conta TP/FN/FP/TN
```

- **Path Resolution:** O projeto é Go com módulo `github.com/rodrigocnascimento/h-search`. Os testes em `tests/` usam `package tests` e importam o pacote público `detector` pelo caminho completo do módulo. Os arquivos de dados são lidos por caminho relativo (`data/benign.txt`, `data/malicious.txt`) a partir de `tests/`.
- **Naming Standards:** Arquivos Go em `snake_case` com sufixo `_test.go`, conforme convenção do projeto. Branch: `refactor/GH-001-ajusta-heuristica-e-testes` (tipo `refactor` por ser ajuste de heurística + reorganização de testes, não nova funcionalidade).

## Validation

- `go test ./...` — todos os testes passam.
- `go vet ./...` — sem avisos.
- Nenhum commit direto em `main`; trabalho realizado na branch `refactor/GH-001-ajusta-heuristica-e-testes`.
