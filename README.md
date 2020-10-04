# ComplexityCounter

A tool to crudely calculate code complexity per file. 

## Assumptions
I'm calculating complexity here assuming code features which result in code branches increase complexity by one point.

Features:
-  if statements
-  ternary conditions
-  logical OR
-  logical AND
-  nullish coalescing
-  switch cases
-  optional chaining

# Flags 

`-extensionsRegex "\\.jsx$"`
- A regex which can be used to change what files you're including. 
- Defaults to `"\\.(js|jsx|ts|jsx)$"`

`-directory ~/Code/Project/src`
- Which directory to run the search on
- Defaults to `"./"`

`-exclusionsRegex __tests__`
- Match on any files you don't want included
- Defaults to `"(__tests__|test\\.(js|ts|jsx|tsx))|__snapshots__"`

`-outputType json`
- Specify `json` or `csv` as an output type
- Defaults to `"csv"`

## JSON Output

```
[
 {
  "filepath": "TestFile.js",
  "lineCount": 50,
  "codeLines": 48,
  "complexity": 7
 }
]
```

## CSV Output

```
file path,line count,code lines,complexity
TestFile.js,50,48,7
```