import { UrlReader } from '@backstage/backend-common';
import { Entity } from '@backstage/catalog-model';
import { Config } from '@backstage/config';
import {
    EntityProvider,
    EntityProviderConnection,
} from '@backstage/plugin-catalog-node';

/**
 * Provides entities from fictional Vela service.
 */
export class VelaProvider implements EntityProvider {
    private readonly env: string;
    private readonly reader: UrlReader;
    private connection?: EntityProviderConnection;
    private hostname: string;

    /** [1] **/
    constructor(env: string, reader: UrlReader, config: Config) {
        this.env = env;
        this.reader = reader;
        this.hostname = config.getOptionalString('vela.host') ?? 'http://localhost/';
    }

    /** [2] **/
    getProviderName(): string {
        return `vela-${this.env}`;
    }

    /** [3] **/
    async connect(connection: EntityProviderConnection): Promise<void> {
        this.connection = connection;
    }

    /** [4] **/
    async run(): Promise<void> {
        if (!this.connection) {
            throw new Error('Not initialized');
        }

        const raw = await this.reader.read(
            this.hostname,
        );
        const data = JSON.parse(raw.toString());

        /** [5] **/
        const entities: Entity[] = data;

        /** [6] **/
        await this.connection.applyMutation({
            type: 'full',
            entities: entities.map(entity => ({
                entity,
                locationKey: `vela-provider:${this.env}`,
            })),
        });
    }
}
