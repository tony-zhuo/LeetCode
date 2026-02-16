import sys


def pytest_collect_file(parent, file_path):
    """Ensure each test file imports solution.py from its own directory."""
    if file_path.suffix == ".py" and file_path.name.startswith("test_"):
        test_dir = str(file_path.parent)
        # Clear cached solution module so it reimports from the correct directory
        for mod_name in list(sys.modules):
            if mod_name == "solution" or mod_name.startswith("solution."):
                del sys.modules[mod_name]
        # Ensure the test directory is first in sys.path
        if test_dir in sys.path:
            sys.path.remove(test_dir)
        sys.path.insert(0, test_dir)
