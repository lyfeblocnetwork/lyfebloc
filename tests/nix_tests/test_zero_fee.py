from pathlib import Path

import pytest
from web3 import Web3

from .network import create_snapshots_dir, setup_custom_lyfebloc
from .utils import (
    ADDRS,
    KEYS,
    eth_to_bech32,
    memiavl_config,
    send_transaction,
    wait_for_fn,
)


@pytest.fixture(scope="module")
def custom_lyfebloc(tmp_path_factory):
    yield from setup_custom_lyfebloc(
        tmp_path_factory.mktemp("zero-fee"),
        26900,
        Path(__file__).parent / "configs/zero-fee.jsonnet",
    )


@pytest.fixture(scope="module")
def custom_lyfebloc_rocksdb(tmp_path_factory):
    path = tmp_path_factory.mktemp("zero-fee-rocksdb")
    yield from setup_custom_lyfebloc(
        path,
        26810,
        memiavl_config(path, "zero-fee"),
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


def test_cosmos_tx(lyfebloc_cluster):
    """
    test basic cosmos transaction works with zero fees
    """
    denom = "alyfe"
    cli = lyfebloc_cluster.cosmos_cli()
    sender = eth_to_bech32(ADDRS["signer1"])
    receiver = eth_to_bech32(ADDRS["signer2"])
    amt = 1000

    old_src_balance = cli.balance(sender, denom)
    old_dst_balance = cli.balance(receiver, denom)

    tx = cli.transfer(
        sender,
        receiver,
        f"{amt}{denom}",
        gas_prices=f"0{denom}",
        generate_only=True,
    )

    tx = cli.sign_tx_json(tx, sender, max_priority_price=0)

    rsp = cli.broadcast_tx_json(tx, broadcast_mode="sync")
    assert rsp["code"] == 0, rsp["raw_log"]

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = cli.balance(receiver, denom)
        return old_dst_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)
    assert old_dst_balance + amt == new_dst_balance
    new_src_balance = cli.balance(sender, denom)
    # no fees paid, so sender balance should be
    # initial_balance - amount_sent
    assert old_src_balance - amt == new_src_balance


def test_eth_tx(lyfebloc_cluster):
    """
    test basic Ethereum transaction works with zero fees
    """
    w3: Web3 = lyfebloc_cluster.w3

    sender = ADDRS["signer1"]
    receiver = ADDRS["signer2"]
    amt = 1000

    old_src_balance = w3.eth.get_balance(sender)
    old_dst_balance = w3.eth.get_balance(receiver)

    receipt = send_transaction(
        w3,
        {
            "from": sender,
            "to": receiver,
            "value": amt,
            "gasPrice": 0,
        },
        KEYS["signer1"],
    )
    assert receipt.status == 1

    new_dst_balance = 0

    def check_balance_change():
        nonlocal new_dst_balance
        new_dst_balance = w3.eth.get_balance(receiver)
        return old_dst_balance != new_dst_balance

    wait_for_fn("balance change", check_balance_change)
    assert old_dst_balance + amt == new_dst_balance
    new_src_balance = w3.eth.get_balance(sender)
    # no fees paid, so sender balance should be
    # initial_balance - amount_sent
    assert old_src_balance - amt == new_src_balance