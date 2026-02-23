# Go Math Expression Parser

## Purpose

Go Math Expression Parser evaluates mathematical expressions for rule conditions and calculations. It enables dynamic rule logic evaluation without code recompilation.

## Type

Shared library

## Consumed By

- rules-service
- strategy-service
- backtest-manager-service

## Exposed Interface

**Key Functions:**

- `Parse(expression string) (Evaluator, error)`: Compile expression
- `Evaluate(expr Evaluator, context map[string]float64) (float64, error)`: Execute expression
- `EvaluateBool(expr string, context map[string]float64) (bool, error)`: Boolean expression

**Supported Operators**: `+`, `-`, `*`, `/`, `%`, `^`, `==`, `!=`, `<`, `>`, `<=`, `>=`, `&&`, `||`

## Usage Example

```go
import "github.com/arconomy/go-math-expression-parser"

evaluator, _ := parser.Parse("price > sma20 && rsi < 30")
result, _ := parser.EvaluateBool(evaluator, map[string]float64{
    "price": 100.5,
    "sma20": 99.0,
    "rsi": 25.0,
})
```

## MCP Rules

**When This Library Changes, Services That Must Be Updated**:
- rules-service (primary consumer)
- Any service using rule evaluation

---

**Last Updated**: 2026-02-22



After ANY change to this service that affects:
- API surface (new/modified endpoints or gRPC methods)
- NATS subjects (new publish/subscribe calls)
- Data models (schema/migration changes)
- Config vars (new env vars)
- Inter-service dependencies (new gRPC clients or NATS consumers)

You MUST update this service's AGENTS.md to reflect the change before committing.
Update the root ./AGENTS.md if the change affects the architecture graph,
NATS subject registry, or gRPC relationship table.

---

**Last Updated**: 2026-02-23
