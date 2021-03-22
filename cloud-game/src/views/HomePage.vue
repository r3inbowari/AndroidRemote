<template>
  <el-container>
    <el-header>
      <div id="jjjj">
        <div>as</div>
        <div>as</div>
        <div>as</div>
        <div>as</div>
      </div>
    </el-header>

    <el-container>
      <el-main
        ><div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div>
        <div>a</div></el-main
      >
    </el-container>
  </el-container>
</template>

<style>
#jjjj {
  box-shadow: 0 2px 4px 0 rgb(0 0 0 / 8%);
}
.el-header,
.el-footer {
  background-color: rgb(21, 21, 21);

  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: rgb(28, 32, 34);
  color: #333;
  text-align: center;
  line-height: 200px;
}

.el-main {
  background-color: rgb(255, 255, 255);
  color: #333;
  text-align: center;
  line-height: 160px;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>
