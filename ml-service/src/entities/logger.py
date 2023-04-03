import logging

from io import TextIOWrapper


# logging
def setup_default_logger(name, log_file, level=logging.INFO, mode="w"):
    """To setup as many loggers as you want"""

    formatter = logging.Formatter("%(asctime)s %(levelname)s %(message)s")

    if type(log_file) is TextIOWrapper:
        handler = logging.StreamHandler(log_file)
    else:
        handler = logging.FileHandler(log_file, mode=mode)
    handler.setFormatter(formatter)

    logger = logging.getLogger(name)
    logger.setLevel(level)
    logger.addHandler(handler)

    return logger
