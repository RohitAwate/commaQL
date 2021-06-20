// Copyright 2021 Rohit Awate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ast

type Node interface {
	// Have to resort to shit like this thanks to lack of a strict inheritance system in Go.
	// Translates to "Hey, I am a Node!" <sigh>
	//
	// Inspired by Go's own AST implementation.
	// https://github.com/golang/go/blob/e3cb3817049ca5e9d96543500b72117f6ca659b8/src/go/ast/ast.go#L33-L36
	// https://github.com/golang/go/blob/e3cb3817049ca5e9d96543500b72117f6ca659b8/src/go/ast/ast.go#L534-L559
	amNode()
}

type Stmt interface {
	Node
	amStmt()
}

type Expr interface {
	Node
	amExpr()
}

type Clause interface {
	Node
	amClause()
}
