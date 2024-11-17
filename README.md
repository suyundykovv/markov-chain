# Markov Chain Text Generator

## Overview

This program generates text based on the **Markov Chain algorithm**. The text is generated using an input source (a text file) and can be customized using options such as the maximum number of words, a starting prefix, and the prefix length.

By default, the program reads the input text from `stdin`, applies the Markov Chain algorithm, and outputs the generated text.

## Features

### Baseline Functionality

- **Default Prefix Length**: 2 words.
- **Suffix Length**: Always 1 word.
- **Maximum Words**: Defaults to 100 words.

### Command-Line Options

1. **Maximum Number of Words (`-w N`)**  
   Generates up to `N` words. The number cannot be negative and cannot exceed 10,000.
   - Default: 100 words.

2. **Starting Prefix (`-p S`)**  
   Starts the generated text with the specified prefix `S`. The prefix must be present in the original text.
   
3. **Prefix Length (`-l N`)**  
   Sets the number of words in the starting prefix. The value must be between 1 and 5.
   - Default: 2 words.

4. **Help (`--help`)**  
   Prints the usage information.

### Error Handling
- Errors are printed with clear messages, e.g., missing input text or invalid parameters (e.g., negative numbers, exceeding limits).
  
### Example Usage

#### 1. Default Behavior (Generate 100 Words)
By default, the program will read from `stdin` and generate 100 words based on the Markov Chain algorithm.

```bash
$ cat the_great_gatsby.txt | ./markovchain | cat -e
Chapter 1 In my younger and more stable, become for a job. He hadn't eat anything for a long, silent time. It was the sound of someone splashing after us over the confusion a long many-windowed room which overhung the terrace. Eluding Jordan's undergraduate who was well over sixty, and Maurice A. Flink and the great bursts of leaves growing on the air now. "How do you want? What do you like Europe?" she exclaimed surprisingly. "I just got here a minute. "Yes." He hesitated. "Was she killed?" "Yes." "I thought you didn't, if you'll pardon my--you see, I carry$
```

#### 2. Limit the Number of Words (`-w N`)
Specify the maximum number of words to generate with the `-w` flag.

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 | cat -e
Chapter 1 In my younger and more stable, become for$
```

#### 3. Specify a Starting Prefix (`-p S`)
The generated text can be forced to start with a given prefix.

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to play" | cat -e
to play for you in that vast obscurity beyond the$
```

#### 4. Specify Prefix Length (`-l N`)
You can set the number of words in the prefix that the Markov Chain algorithm should use.

```bash
$ cat the_great_gatsby.txt | ./markovchain -w 10 -p "to something funny" -l 3
to something funny the last two days," remarked Wilson. "That's$
```

#### 5. Help Option (`--help`)
Display the usage instructions.

```bash
$ ./markovchain --help
Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words
  -p S    Starting prefix
  -l N    Prefix length
```

---

## Constraints

- **Maximum Number of Words (`-w N`)**:
  - Must be **non-negative** and not exceed 10,000.
  
- **Starting Prefix (`-p S`)**:
  - The provided prefix must **exist in the original text**.

- **Prefix Length (`-l N`)**:
  - Must be between **1 and 5** (inclusive).
  
- **Error Handling**:
  - Missing or invalid input will result in an error message indicating the issue, such as:
    - No input text provided (`Error: no input text`)
    - Invalid number of words (`Error: invalid number of words`)
    - Invalid prefix length (`Error: invalid prefix length`)

---

## Example Workflow

1. Prepare your input text (e.g., from a file).
2. Pipe the input into the program using `cat` or another method.
3. Use the appropriate flags to control the output:
   - `-w N`: Set maximum words.
   - `-p S`: Set the starting prefix.
   - `-l N`: Set the prefix length.
4. Review the generated text and adjust parameters as needed.

---

## Error Handling and Validation

### Common Errors:

1. **Missing Input**:
   If no input text is provided, the program will output:
   ```bash
   Error: no input text
   ```

2. **Invalid Number of Words (`-w N`)**:
   If the `-w` value is negative or exceeds 10,000, the program will output:
   ```bash
   Error: invalid number of words
   ```

3. **Prefix Not Found**:
   If the specified prefix `-p S` does not exist in the input text, the program will output:
   ```bash
   Error: prefix not found
   ```

4. **Invalid Prefix Length (`-l N`)**:
   If the prefix length is outside the valid range (1 to 5), the program will output:
   ```bash
   Error: invalid prefix length
   ```

---

## Requirements

- **Input Source**: The program reads from `stdin`. You can pipe text from a file or other command.
- **Programming Language**: This project is implemented in a programming language like Python, C, or C++ (depending on the implementation).

---

## Conclusion

This Markov Chain text generator creates random text based on a given source text. It allows for flexibility in the number of words generated, the starting prefix, and the prefix length. With error handling and usage options, it provides an easy way to generate creative, random text from existing content.