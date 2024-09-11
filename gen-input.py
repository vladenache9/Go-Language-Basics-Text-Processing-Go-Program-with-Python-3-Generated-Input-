import os
import random
import string
import sys
import hashlib

def generate_matrix(rows, cols):
    return ''.join(random.choice('01') for _ in range(rows * cols))

def bin_to_hex(binary_str):
    hex_str = ''
    for i in range(0, len(binary_str), 4):
        nibble = binary_str[i:i+4]
        hex_str += format(int(nibble, 2), 'X')
    return hex_str

def generate_in_file(file_path, min_lines, duplicate_percentage, use_cache):
    with open(file_path, 'w') as f:
        lines = []
        cache = set() if use_cache else None  
        duplicate_lines = int(min_lines * duplicate_percentage / 100)
        unique_lines = min_lines - duplicate_lines
        
       
        for _ in range(unique_lines):
            rows = random.randint(10, 10)  
            cols = random.randint(10, 10)  
            matrix = generate_matrix(rows, cols)
            line = f"{rows}x{cols}:{matrix}"
            
            if use_cache and line in cache:
                continue  
            
            lines.append(line)
            if use_cache:
                cache.add(line)  
        
      
        for _ in range(duplicate_lines):
            lines.append(random.choice(lines))
        
       
        random.shuffle(lines)
        
       
        for line in lines:
            f.write(line + '\n')

def generate_x_file(file_path, in_file_path, use_cache):
    cache = set() if use_cache else None  
    with open(in_file_path, 'r') as in_file, open(file_path, 'w') as out_file:
        for line in in_file:
            matrix_info, binary_data = line.strip().split(':')
            if use_cache and line in cache:
                continue  
            hex_data = bin_to_hex(binary_data)
            out_file.write(f"{matrix_info}:{hex_data}\n")
            if use_cache:
                cache.add(line) 

def main(output_dir, file_type, num_files, min_lines, duplicate_percentage, use_cache):
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    for i in range(num_files):
        file_name = f"mat_{i+1}.in"
        file_path = os.path.join(output_dir, file_name)
        
      
        generate_in_file(file_path, min_lines, duplicate_percentage, use_cache)
        
        if file_type == "x":
           
            x_file_path = file_path + ".x"
            generate_x_file(x_file_path, file_path, use_cache)
            os.remove(file_path) 

if __name__ == "__main__":
    if len(sys.argv) < 6 or len(sys.argv) > 7:
        print("Usage: python3 gen-input.py <output-dir> <in/x> <num-files> <min-lines> <duplicate-percentage> [nocache]")
        sys.exit(1)
    
    output_dir = sys.argv[1]
    file_type = sys.argv[2]
    num_files = int(sys.argv[3])
    min_lines = int(sys.argv[4])
    duplicate_percentage = float(sys.argv[5].rstrip('%'))
    
    use_cache = True
    if len(sys.argv) == 7 and sys.argv[6] == "nocache":
        use_cache = False  
    
    if file_type not in ["in", "x"]:
        print("File type must be 'in' or 'x'.")
        sys.exit(1)
    
    if not (0 <= duplicate_percentage <= 100):
        print("Duplicate percentage must be between 0 and 100.")
        sys.exit(1)
    
    main(output_dir, file_type, num_files, min_lines, duplicate_percentage, use_cache)
