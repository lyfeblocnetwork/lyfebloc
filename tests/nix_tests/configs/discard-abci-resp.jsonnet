local config = import 'default.jsonnet';

config {
  'lyfebloc_1775-1'+: {
    config+: {
      storage: {
        discard_abci_responses: true,
      },
    },
  },
}
