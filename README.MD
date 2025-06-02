# tree-sitter-onescript

**tree-sitter-onescript** — полноценная грамматика [Tree-sitter](https://tree-sitter.github.io/), для языка [OneScript](https://github.com/EvilBeaver/OneScript).

---

## Быстрый старт

```bash
# 0. Требования
#    ───────────
#    • Rust ≥ 1.77 (edition 2024)  • tree-sitter-cli ≥ 0.25  • Node.js — по желанию

# 1. Подключаем зависимость
cargo add tree-sitter-onescript      # или path-зависимость, пока не выложено на crates.io

# 2. Мини-пример на Rust
use tree_sitter::{Parser};
use tree_sitter_onescript::language;

let mut parser = Parser::new();
parser.set_language(language()).unwrap();
let tree = parser.parse("Перем x; x = 1;", None).unwrap();
println!("{}", tree.root_node().to_sexp());
```

# Сборка из исходников

```bash
git clone https://github.com/tree-sitter-onescript
cd tree-sitter-onescript
tree-sitter generate   # генерирует src/parser.c
cargo build            # собирает rlib + cdylib

```

# Структура проекта

```bash
├── grammar.js          # грамматика
├── src/
│   ├── parser.c        # автогенерируемый C-парсер
│   └── lib.rs          # безопасная обёртка (pub fn language())
└── test/corpus         # corpus-тесты (tree-sitter test …)

```

# TODO

* `highlights.scm`

* Автоматическое тестирование на популярных библиотеках OneScript