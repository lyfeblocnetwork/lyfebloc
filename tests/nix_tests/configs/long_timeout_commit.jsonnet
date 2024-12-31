local default = import 'default.jsonnet';

default {
  'lyfebloc_1775-1'+: {
    config+: {
      consensus+: {
        timeout_commit: '5s',
      },
    },
  },
}
