# config/logging_config.py
import logging
import logging.config
import os

LOGGING_CONFIG = {
    "version": 1,
    "disable_existing_loggers": False,
    "formatters": {
        "standard": {
            "format": "%(asctime)s [%(levelname)s] %(name)s: %(message)s"
        },
    },
    "handlers": {
        "console": {
            "class": "logging.StreamHandler",
            "formatter": "standard",
        },
        "file": {
            "class": "logging.FileHandler",
            "filename": os.path.join("logs", "project.log"),
            "formatter": "standard",
        },
    },
    "root": {
        "handlers": ["console", "file"],
        "level": "INFO",
    },
}

def setup_logging():
    logging.config.dictConfig(LOGGING_CONFIG)
