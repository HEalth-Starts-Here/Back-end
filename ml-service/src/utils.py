import gdown

from src.entities.predict_params import DownloadParams


def download_file(params: DownloadParams):
    gdown.download(params.file_url, params.path_to_save, fuzzy=True)
