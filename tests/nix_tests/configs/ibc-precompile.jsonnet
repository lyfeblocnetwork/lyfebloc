local config = import 'ibc.jsonnet';

config {
  'lyfebloc_1775-1'+: {
    genesis+: {
      app_state+: {
        feemarket+: {
          params+: {
            no_base_fee: false,
            base_fee: '100000000000',
          },
        },
      },
    },
  },
}
