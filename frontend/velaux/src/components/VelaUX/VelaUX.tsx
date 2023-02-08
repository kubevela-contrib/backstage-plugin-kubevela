import React, { useEffect, useState } from 'react';
import { styled, TextField, OutlinedInputProps } from '@material-ui/core';
import {
  Header,
  Page,
} from '@backstage/core-components';

const defaultVelaUxUrl = 'http://localhost:9082';

const CssTextField = styled(TextField)({
  '& label.Mui-focused': {
    color: '#333333',
  },
});

export const VelaUX = () => {
  const [velaUxUrl, setVelaUxUrl] = useState<string>(defaultVelaUxUrl)

  useEffect(() => {
    const storedUrl = localStorage.getItem('velaUxUrl')
    if (storedUrl) {
      setVelaUxUrl(storedUrl)
    }
  }, [])

  return (
    <Page themeId="tool">
      <Header
        title="Welcome VelaUX!"
        subtitle="KubeVela User Experience (UX). An extensible, application-oriented delivery and management Dashboard. Specify your VelaUX endpoint on the right."
      >
        <CssTextField
          style={{ width: '100%', padding: 0, background: "#ffffff80", borderRadius: 6, backdropFilter: "blur(16px)" }}
          InputProps={{ disableUnderline: true } as Partial<OutlinedInputProps>}
          required
          variant="filled"
          color="primary"
          defaultValue={velaUxUrl}
          value={velaUxUrl}
          label='VelaUX Endpoint'
          onChange={e => {
            setVelaUxUrl(e.target.value)
            localStorage.setItem('velaUxUrl', e.target.value)
          }}
        ></CssTextField>
      </Header>
      <iframe
        id="velaux-iframe"
        style={{
          border: 0,
          width: "100%",
          height: "100%",
          gridArea: "pageContent",
        }}
        src={velaUxUrl}
      />
    </Page >
  )
};
