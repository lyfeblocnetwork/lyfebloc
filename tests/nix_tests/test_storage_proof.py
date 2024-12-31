import pytest

from .network import setup_lyfebloc, setup_lyfebloc_rocksdb
from .utils import CONTRACTS, deploy_contract, w3_wait_for_new_blocks


@pytest.fixture(scope="module")
def custom_lyfebloc(tmp_path_factory):
    path = tmp_path_factory.mktemp("storage-proof")
    yield from setup_lyfebloc(path, 26800)


@pytest.fixture(scope="module")
def custom_lyfebloc_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("storage-proof-rocksdb")
    yield from setup_lyfebloc_rocksdb(path, 26810)


@pytest.fixture(scope="module", params=["lyfebloc", "lyfebloc-rocksdb", "geth"])
def cluster(request, custom_lyfebloc, custom_lyfebloc_rocksdb, geth):
    """
    run on both lyfebloc (default build and rocksdb)
    and geth
    """
    provider = request.param
    if provider == "lyfebloc":
        yield custom_lyfebloc
    elif provider == "lyfebloc-rocksdb":
        yield custom_lyfebloc_rocksdb
    elif provider == "geth":
        yield geth
    else:
        raise NotImplementedError


def test_basic(cluster):
    # wait till height > 2 because
    # proof queries at height <= 2 are not supported
    if cluster.w3.eth.block_number <= 2:
        w3_wait_for_new_blocks(cluster.w3, 2)

    _, res = deploy_contract(
        cluster.w3,
        CONTRACTS["StateContract"],
    )
    method = "eth_getProof"
    storage_keys = ["0x0", "0x1"]
    proof = (
        cluster.w3.provider.make_request(
            method, [res["contractAddress"], storage_keys, hex(res["blockNumber"])]
        )
    )["result"]
    for proof in proof["storageProof"]:
        if proof["key"] == storage_keys[0]:
            assert proof["value"] != "0x0"
