package gogoproto

import "testing"

func TestCaseConvert(t *testing.T) {
	camelCase := "helloWorld"
	pascalCase := "HelloWorld"
	snakeCase := "hello_world"

	// 测试从 CamelCase 到 PascalCase 和 SnakeCase 的转换
	t.Log(ToCamelCase(camelCase))  // 输出："helloWorld"
	t.Log(ToCamelCase(snakeCase))  // 输出："helloWorld"
	t.Log(ToCamelCase(pascalCase)) // 输出："helloWorld"

	t.Log(ToPascalCase(camelCase))  // 输出："HelloWorld"
	t.Log(ToPascalCase(pascalCase)) // 输出："HelloWorld"
	t.Log(ToPascalCase(snakeCase))  // 输出："HelloWorld"

	t.Log(ToSnakeCase(camelCase))  // 输出："hello_world"
	t.Log(ToSnakeCase(pascalCase)) // 输出："hello_world"
	t.Log(ToSnakeCase(snakeCase))  // 输出："hello_world"
}

func TestCaseConvert1(t *testing.T) {
	camelCase := "helloWorld1"
	pascalCase := "HelloWorld1"
	snakeCase := "hello_world1"
	snakeCase1 := "hello_world_1"

	// 测试从 CamelCase 到 PascalCase 和 SnakeCase 的转换
	t.Log(ToCamelCase(camelCase))  // 输出："helloWorld1"
	t.Log(ToCamelCase(snakeCase))  // 输出："helloWorld1"
	t.Log(ToCamelCase(pascalCase)) // 输出："helloWorld1"
	t.Log(ToCamelCase(snakeCase1))

	t.Log(ToPascalCase(camelCase))  // 输出："HelloWorld1"
	t.Log(ToPascalCase(pascalCase)) // 输出："HelloWorld1"
	t.Log(ToPascalCase(snakeCase))  // 输出："HelloWorld1"
	t.Log(ToPascalCase(snakeCase1))

	t.Log(ToSnakeCase(camelCase))  // 输出："hello_world1"
	t.Log(ToSnakeCase(pascalCase)) // 输出："hello_world1"
	t.Log(ToSnakeCase(snakeCase))  // 输出："hello_world1"
	t.Log(ToSnakeCase(snakeCase1))
}
