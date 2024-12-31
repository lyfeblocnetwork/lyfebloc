import pytest

from .network import setup_lyfebloc, setup_lyfebloc_rocksdb, setup_geth


@pytest.fixture(scope="session")
def lyfebloc(tmp_path_factory):
    path = tmp_path_factory.mktemp("lyfebloc")
    yield from setup_lyfebloc(path, 26650)


@pytest.fixture(scope="session")
def lyfebloc_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("lyfebloc-rocksdb")
    yield from setup_lyfebloc_rocksdb(path, 20650)


@pytest.fixture(scope="session")
def geth(tmp_path_factory):
    path = tmp_path_factory.mktemp("geth")
    yield from setup_geth(path, 8545)


@pytest.fixture(scope="session", params=["lyfebloc", "lyfebloc-ws"])
def lyfebloc_rpc_ws(request, lyfebloc):
    """
    run on both lyfebloc and lyfebloc websocket
    """
    provider = request.param
    if provider == "lyfebloc":
        yield lyfebloc
    elif provider == "lyfebloc-ws":
        lyfebloc_ws = lyfebloc.copy()
        lyfebloc_ws.use_websocket()
        yield lyfebloc_ws
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["lyfebloc", "lyfebloc-ws", "lyfebloc-rocksdb", "geth"])
def cluster(request, lyfebloc, lyfebloc_rocksdb, geth):
    """
    run on lyfebloc, lyfebloc websocket,
    lyfebloc built with rocksdb (memIAVL + versionDB)
    and geth
    """
    provider = request.param
    if provider == "lyfebloc":
        yield lyfebloc
    elif provider == "lyfebloc-ws":
        lyfebloc_ws = lyfebloc.copy()
        lyfebloc_ws.use_websocket()
        yield lyfebloc_ws
    elif provider == "geth":
        yield geth
    elif provider == "lyfebloc-rocksdb":
        yield lyfebloc_rocksdb
    else:
        raise NotImplementedError


@pytest.fixture(scope="module", params=["lyfebloc", "lyfebloc-rocksdb"])
def lyfebloc_cluster(request, lyfebloc, lyfebloc_rocksdb):
    """
    run on lyfebloc default build &
    lyfebloc with rocksdb build and memIAVL + versionDB
    """
    provider = request.param
    if provider == "lyfebloc":
        yield lyfebloc
    elif provider == "lyfebloc-rocksdb":
        yield lyfebloc_rocksdb
    else:
        raise NotImplementedError
