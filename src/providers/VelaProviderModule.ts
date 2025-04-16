import {
  coreServices,
  createBackendModule,
  readSchedulerServiceTaskScheduleDefinitionFromConfig,
} from '@backstage/backend-plugin-api';
import { VelaProvider } from './VelaProvider';
import { catalogProcessingExtensionPoint } from '@backstage/plugin-catalog-node/alpha';

export const velaProviderModule = createBackendModule({
  pluginId: 'catalog',
  moduleId: 'vela-provider',
  register(env) {
    env.registerInit({
      deps: {
        catalog: catalogProcessingExtensionPoint,
        reader: coreServices.urlReader,
        scheduler: coreServices.scheduler,
        rootConfig: coreServices.rootConfig
      },
      async init({ catalog, reader, scheduler, rootConfig }) {
        const config = rootConfig.getConfig('catalog.providers.vela')
        const schedule = config.has('schedule')
          ? readSchedulerServiceTaskScheduleDefinitionFromConfig(
              config.getConfig('schedule'),
            )
          : {
              frequency: { seconds: 60 },
              timeout: { seconds: 600 },
            };
        const taskRunner = scheduler.createScheduledTaskRunner(schedule);
        const vela = new VelaProvider('dev', reader, config);
        catalog.addEntityProvider(vela);
        taskRunner.run({
          id: 'vela-provider-refresh',
          fn: async () => {
            await vela.run();
          },
        });
      },
    });
  },
});
