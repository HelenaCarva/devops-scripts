import argparse
import logging
import os
import sys

# Set up logging configuration
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def parse_arguments():
    parser = argparse.ArgumentParser(description='DevOps Scripts')
    parser.add_argument('--script', type=str, required=True, help='Script to execute')
    parser.add_argument('--args', type=str, nargs='+', help='Arguments to pass to the script')
    return parser.parse_args()

def execute_script(script, args):
    try:
        module = __import__(script)
        if hasattr(module, 'main'):
            module.main(args)
        else:
            logging.error(f"Script '{script}' does not have a 'main' function")
            sys.exit(1)
    except ImportError:
        logging.error(f"Script '{script}' not found")
        sys.exit(1)

def main():
    args = parse_arguments()
    script = args.script
    script_args = args.args
    execute_script(script, script_args)

if __name__ == '__main__':
    main()