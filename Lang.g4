grammar Lang;

program: namespace (nClass | nInterface)* EOF;

namespace: 'package' PACKAGE_NAME ';';

annotation: '@' identifier '(' key_value? (',' key_value)* ')';

key_value: (identifier '=' value);

modifier: 'private' | 'public';

nInterface: modifier? 'interface' identifier '{' '}';

nClass:
	annotation* modifier? 'class' identifier '{' classContent* '}';

classContent: fieldDeclaration | methodDeclaration;

fieldDeclaration: modifier? langType identifier ';';

methodDeclaration:
	annotation* modifier? identifier '(' ')' '{' '}';

identifier: TYPE_NAME | PACKAGE_NAME;

langType: 'String' | 'Int' | 'Bool';

value: INTEGER;

INTEGER: [0-9]+;

PACKAGE_NAME: [a-z]+;

TYPE_NAME: [A-Za-z] [A-Za-z0-9]*;

WS: [ \t\r\n]+ -> skip;
