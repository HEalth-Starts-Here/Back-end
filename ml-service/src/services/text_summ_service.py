import sys

import spacy
import pytextrank

from protos.text_summarization_pb2 import TextSummResponse
from protos.text_summarization_pb2_grpc import TextSummServicer
from src.entities.logger import setup_default_logger
from src.entities.models_params import TextSummParams


logger = setup_default_logger("text_summ_logs", sys.stdout)


class TextSummService(TextSummServicer):
    def __init__(self, params: TextSummParams) -> None:
        self.params = params

    def summarizeText(self, request, context):
        raw_txt = request.text

        nlp = spacy.load(self.params.nlp_sum_type)
        nlp.add_pipe("textrank", last=True)
        doc = nlp(raw_txt)

        summarized_text = []
        for sent in doc._.textrank.summary(limit_sentences=self.params.limit_sentence):
            summarized_text.append(str(sent))
        summarized_text = "".join(summarized_text)
        logger.info(f"Summarized text lenght: {len(summarized_text)}")
        return TextSummResponse(summarized_text=summarized_text)
