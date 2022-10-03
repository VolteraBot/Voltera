<script>
  import {
    Header,
    HeaderUtilities,
    HeaderGlobalAction,
    SideNav,
    SideNavItems,
    SideNavMenu,
    SideNavMenuItem,
    SideNavLink,
    SkipToContent,
    Content,
    Grid,
    Row,
    Column,
  } from "carbon-components-svelte";

  import Fade from "carbon-icons-svelte/lib/Fade.svelte"
  import VolteraIcon from "./components/VolteraIcon.svelte";
  import VolteraTitle from "./components/VolteraTitle.svelte";

  let isSideNavOpen = false;
</script>

<Header bind:isSideNavOpen data-tauri-drag-region>
  <svelte:fragment slot="platform">
    <VolteraIcon />
    <VolteraTitle />
  </svelte:fragment>
  <svelte:fragment slot="skip-to-content">
    <SkipToContent />
  </svelte:fragment>
  <HeaderUtilities>
    <!--<HeaderGlobalAction aria-label="Minimize" icon={Subtract} on:click={appWindow.minimize()} />
    <HeaderGlobalAction aria-label="Maximize" icon={CollapseAll} on:click={appWindow.toggleMaximize()}/>
    <HeaderGlobalAction aria-label="Close" icon={Close} on:click={appWindow.close()}/>-->
  </HeaderUtilities>
</Header>
<SideNav bind:isOpen={isSideNavOpen} rail>
  <SideNavItems>
    <SideNavLink icon={Fade} text="Link 1" />
    <SideNavLink icon={Fade} text="Link 2" />
    <SideNavLink icon={Fade} text="Link 3" />
    <SideNavMenu icon={Fade} text="Menu">
      <SideNavMenuItem href="/" text="Link 1" />
      <SideNavMenuItem href="/" text="Link 2" />
      <SideNavMenuItem href="/" text="Link 3" />
    </SideNavMenu>
  </SideNavItems>
</SideNav>
<Content>
  <Grid>
    <Row>
      <Column>
        <h1>Welcome</h1>
      </Column>
    </Row>
  </Grid>
</Content>