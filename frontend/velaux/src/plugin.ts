import { createPlugin, createRoutableExtension } from '@backstage/core-plugin-api';

import { rootRouteRef } from './routes';

export const velauxPlugin = createPlugin({
  id: 'velaux',
  routes: {
    root: rootRouteRef,
  },
});

export const VelauxPage = velauxPlugin.provide(
  createRoutableExtension({
    name: 'VelauxPage',
    component: () =>
      import('./components/VelaUX').then(m => m.VelaUX),
    mountPoint: rootRouteRef,
  }),
);
