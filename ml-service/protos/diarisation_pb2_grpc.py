# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import protos.diarisation_pb2 as diarisation__pb2


class DiarisationStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.transcribeAudio = channel.unary_unary(
                '/Diarisation/transcribeAudio',
                request_serializer=diarisation__pb2.DiarisationRequest.SerializeToString,
                response_deserializer=diarisation__pb2.DiarisationResponse.FromString,
                )


class DiarisationServicer(object):
    """Missing associated documentation comment in .proto file."""

    def transcribeAudio(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DiarisationServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'transcribeAudio': grpc.unary_unary_rpc_method_handler(
                    servicer.transcribeAudio,
                    request_deserializer=diarisation__pb2.DiarisationRequest.FromString,
                    response_serializer=diarisation__pb2.DiarisationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Diarisation', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Diarisation(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def transcribeAudio(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Diarisation/transcribeAudio',
            diarisation__pb2.DiarisationRequest.SerializeToString,
            diarisation__pb2.DiarisationResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
