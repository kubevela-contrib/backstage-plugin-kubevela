import {
  coreServices,
  createBackendModule,
  SchedulerServiceTaskScheduleDefinition,
  readSchedulerServiceTaskScheduleDefinitionFromConfig,
} from '@backstage/backend-plugin-api';
import { VelaProvider } from './VelaProvider';

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
        const frequency: number = config.getOptionalNumber('frequency') || 60;
        const timeout: number = config.getOptionalNumber('timeout') || 600;
        const hostname:string = config.getConfig('host');
        const schedule = config.has('schedule')
          ? readSchedulerServiceTaskScheduleDefinitionFromConfig(
              config.getConfig('schedule'),
            )
          : {
              frequency: { seconds: 60 },
              timeout: { seconds: 600 },
            };
        const taskRunner = scheduler.createScheduledTaskRunner(schedule);
        const vela = new VelaProvider('dev', reader, taskRunner, hostname);
        catalog.addEntityProvider(vela);
      },
    });
  },
});
