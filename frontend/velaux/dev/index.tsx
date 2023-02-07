import React from 'react';
import { createDevApp } from '@backstage/dev-utils';
import { velauxPlugin, VelauxPage } from '../src/plugin';

createDevApp()
  .registerPlugin(velauxPlugin)
  .addPage({
    element: <VelauxPage />,
    title: 'Root Page',
    path: '/velaux'
  })
  .render();
