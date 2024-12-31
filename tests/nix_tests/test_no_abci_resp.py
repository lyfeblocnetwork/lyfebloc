from pathlib import Path

import pytest

from .network import create_snapshots_dir, setup_custom_lyfebloc
from .utils import memiavl_config, wait_for_block


@pytest.fixture(scope="module")
def custom_lyfebloc(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp")
    yield from setup_custom_lyfebloc(
        path,
        26260,
        Path(__file__).parent / "configs/discard-abci-resp.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_lyfebloc_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("no-abci-resp-rocksdb")
    yield from setup_custom_lyfebloc(
        path,
        26810,
        memiavl_config(path, "discard-abci-resp"),
        post_init=create_snapshots_dir,
        chain_binary="lyfeblocd-rocksdb",
    )


@pytest.fixture(scope="module", params=["lyfebloc", "lyfebloc-rocksdb"])
def lyfebloc_cluster(request, custom_lyfebloc, custom_lyfebloc_rocksdb):
    """
    run on lyfebloc and
    lyfebloc built with rocksdb (memIAVL + versionDB)
    """
    provider = request.param
    if provider == "lyfebloc":
        yield custom_lyfebloc

    elif provider == "lyfebloc-rocksdb":
        yield custom_lyfebloc_rocksdb

    else:
        raise NotImplementedError


def test_gas_eth_tx(lyfebloc_cluster):
    """
    When node does not persist ABCI responses
    eth_gasPrice should return an error instead of crashing
    """
    wait_for_block(lyfebloc_cluster.cosmos_cli(), 3)
    try:
        lyfebloc_cluster.w3.eth.gas_price  # pylint: disable=pointless-statement
        raise Exception(  # pylint: disable=broad-exception-raised
            "This query should have failed"
        )
    except Exception as error:
        assert "block result not found" in error.args[0]["message"]
