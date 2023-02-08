# velaux

This plugin adds a sidebar page to allow you to quickly access VelaUX.

Edit `packages/app/src/components/Root/Root.tsx` to add the sidebar

```diff
 // packages/app/src/components/Root/Root.tsx
+import { VelaLogo } from '@oamdev/backstage-plugin-velaux';

 export const Root = ({ children }: PropsWithChildren<{}>) => (
   <SidebarPage>
         {/* ... */}
         {/* End global nav */}
         <SidebarDivider />
         <SidebarScrollWrapper>
+          <SidebarItem icon={VelaLogo} to="velaux" text="VelaUX" />
         </SidebarScrollWrapper>
       {/* ... */}
   </SidebarPage>
 );
```

Edit `packages/app/src/App.tsx` to add routers

```diff
 // packages/app/src/App.tsx
+import { VelauxPage } from '@oamdev/backstage-plugin-velaux';

 const routes = (
   <FlatRoutes>
     {/* ... */}
     <Route path="/settings" element={<UserSettingsPage />} />
     <Route path="/catalog-graph" element={<CatalogGraphPage />} />
+    <Route path="/velaux" element={<VelauxPage />} />
     {/* ... */}
   </FlatRoutes>
 );
```

Lastly, you need to make sure you can access the VelaUX webpage. For example, by forwarding the VelaUX addon during development:

```
vela port-forward -n vela-system addon-velaux 9082:80
```
