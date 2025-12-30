# GradMotion 产品使用手册

> [!TIP]
> **更新记录**

- 20250520 更新数据 SDK 文档
- 20250617 更新云开发机，云存储，个人镜像

GradMotion `是专为具身智能机器人算法学习训练打造的利器`

用户聚焦机器人算法创新，其余交给 GradMotion（提供训练环境和算力），上传代码一键启动，记录结果，辅助调优

> [!TIP]

## **平台****官方镜像****环境：**

1. **Isaac GYM**:preview-4 | Python:3.8.20 | PyTorch:2.4.1 | Ubuntu 20.04.5 LTS | Cuda:12.1
2. **IsaacSim:4.2** | IsaacLab:1.4.1 | Python:3.10.14 | PyTorch:2.5.1 | Ubuntu 22.04.3 LTS
3. **IsaacSim:4.5** | IsaacLab:2.0 | Python:3.10.15 | PyTorch:2.5.1 | Ubuntu 22.04.5 LTS
4. **IsaacSim:4.5** | IsaacLab:2.1 | Python:3.10.15 | PyTorch:2.5.1 | Ubuntu 22.04.5 LTS
5. **IsaacSim:5.0** | IsaacLab:2.2 | Python:3.11.13 | PyTorch:2.7 | Ubuntu 22.04.5 LTS
6. **Ubuntu 22.04**
7. **Ubuntu 20.04**

## **平台支持**

1. 训练平台：支持多任务场景（强化学习 + 模仿学习）
2. 云开发机：在线 IDE，支持保存成个人镜像
3. 云桌面_：图形计算机_

# 一、视频教程 📺

## 训练平台

## 云开发机

## 云桌面

新增中...

# 二、图文教程 📚

## 【场景 1】训练平台（强化学习）

### 登录平台（GradMotion）

- `登录网址：``https://spaces.gradmotion.com.cn`
- `登录方式：支持3种，邮箱📮、谷歌账号、GitHub`
- `官网网址：``https://gradmotion.com.cn/`

![](static/YA7ybuPhtomeo3xpRk5cjca9nad.png)

### 查询资产（算力资源）

`2.1 点击右上角账户头像`

`2.2 选择个人中心`

![](static/VleWbWfGgovA7VxxsBZcbbp4nKh.png)

`2.3 查看算力资产`

- `确保账户中有可用算力资源资产`

### 新建项目

- `属于存放训练任务的文件夹`

![](static/LTiJbTihgojSB2xvq6Oc0WfJnUg.png)

### 新建任务

- `进入项目，点击新建任务`

![](static/VDY2bWBLjox22kxLgByc23WVnHf.png)

- `填写任务信息`
  - `任务名称`
  - `任务描述`
  - `任务标签`
  - `选择算力资源（此处显示的是用户账户中可用的算力资产）`
  - `选择环境镜像`
  - `上传代码（本地代码包 或 关联Gitlab）`
  - `填入主启动程序配置`
  - `填入超参路径（可选，非必须）`
  - `填入启动参数`  **必须使用 headless 模式启动**，GradMotion 暂不支持 GUI 界面
  - `点击直接运行`
    ![](static/HnjMbu8kooT7FFx580dcs8wAnlh.png)
    ![](static/TLzVbomMWoRVFpxKY51crYfNnlf.png)
    ![](static/V9qQbYF03oFFrMxosa6c8cC7nah.png)

  #### ‼️ 注意事项

  1. 必须加**--headless**参数
  2. 上传的代码中**不要包含 isaacgym**，平台镜像中已经提供 isaacgym
  3. 上传的压缩包中**建议不要包括本地的模型文件**（logs 目录下保存的.pt 文件），文件占用大量内存，影响项目上传速度，删去不影响训练
  4. gm-run 命令必须在一行命令的开头

### 确认任务状态

![](static/XUsCb2SMKoQevUxFEsJcULF2nId.png)

![](static/SPefbmZSLo9CfpxF7OOc3TQGnJc.png)

![](static/Rn3ybU1Hyo44IIx6MujcROwOnqf.png)

### 查看任务结果

- `支持查看：`

  - `任务详情`
  - `任务代码`
  - `任务超参`
  - `任务环境`
  - `任务日志`
  - `任务图表`
  - `任务模型：可播放仿真训练视频回传`
- `支持添加`

  - `任务笔记`
    ![](static/IwzzbWmu8oUN7xx4acScaUnKnZT.png)

### 对比任务

- `支持多任务对比： 任务概述/代码/超参/图表/环境`，差异化配置会被高亮

![](static/CZPpbPQVTolv1fxHITxccreMnVd.png)

![](static/HIpibPLklorAjIxkpjHcFl2MnBb.png)

### 查询资产消耗

![](static/UJhpbM6R8oJXiJxTedqcv7A6nSe.png)

### 常见问题

#### 找不到文件

python: can't open file xxxxxxxx': [Errno 2] No such file or directory 报错，

**检查启动指令中路径输入是否正确**

**压缩包上传的代码**，检查路径是否从压缩包下第一级文件夹开始输入

**github 拉取的代码**，检查是否输入了完整路径，包括红色框内所有内容

例：该目录对应的启动路径如下

gm-run agibot_x1_train/humanoid/scripts/train.py --task=x1_dh_stand --headless

![](static/AOHIbtCIconA9zx76xNc4eHSnBb.png)

#### 运行任务，报错找不到 module

ModuleNotFoundError: No module named 'xxx 第三方库 xxx'

**解决办法：**找到 setup.py 文件所在目录，**在 setup.py 的 install_require 列表中添加缺失的库**

#### 如果缺失 setup.py 文件，按如下格式新增

![](static/Jm0rbb35OoJKmyxyEtnceqc1nPz.png)

#### **若出现 wandb 相关报错，禁用 wandb**

若有相关启动参数，在启动命令处禁用 wandb，如果没有相关启动参数，将代码中
wandb.init(///YOUR PARAMETERS///) 改为 wandb.init(mode="disabled")

#### 某目标文件实际存在，仍然报错 File not found error

在尝试打开该目标文件的脚本中修改，使用 __**file__** 获取当前脚本所在目录，再拼接相对路径，将目标文件路径改为基于当前脚本文件的绝对路径

### 体验机器人示例工程

机器人示例工程代码库，文档中所有示例工程代码都已经在 GradMotion 平台中跑通

[实践-具身 RL 开源工程](https://paj5uamwttr.feishu.cn/wiki/OdJiwnHdKivhv2kFZk7cI9Svnmb)

## 【场景 2】训练平台（模仿学习-挂载数据集）

1. 在个人存储中，上传所需的数据集文件/文件夹

![](static/ToNEbqYtjoZ0r6xSQACc7cvUnld.png)

1. 在项目代码中读取数据集的部分，将文件路径替换为**个人存储中的绝对路径**

因项目引用数据集的方式不同，修改方法有区别，以下修改供参考

![](static/P9ALbEYaeotCIxxe4ixcnknsn1f.png)

1. 正常新建任务，选择参数，上传代码包（不含数据集）
2. 填入代码中调用的数据集位置目录

![](static/Z9KabB90DoC35uxahQQcAzJdnfc.png)

1. 运行模仿学习训练任务，检查任务状态和日志等。

## 【场景 3】云开发机

### 登录平台（GradMotion）

[https://spaces.gradmotion.com.cn/](https://spaces.gradmotion.com.cn/)

### 新建开发机

填写开发机名称
选择环境镜像
选择算力资源
**填写 SSH 公钥（在本地电脑获取 SSH 公钥）**

![](static/TXofbHs4ZoS1vGxtGMxcj5ESngf.png)

获取 SSH 公钥的操作会因操作系统的不同而有所差异，

以下是主流系统（Windows、Linux/macOS）的详细获取方法：

#### **-**** ****Linux/macOS 系统**

1. 检查是否已有公钥

打开终端（Terminal），输入以下命令查看 `~/.ssh` 目录下的文件：

`ls`` ``-la`` ~/.ssh`

若存在 `id_rsa.pub` 或 `id_ed25519.pub` 文件，该文件即为 SSH 公钥。

若没有相关文件，需手动生成新的公钥。

![](static/XQOabA9MBoteDUxOiiBcd0Lbnsc.png)

2. 生成新的 SSH 公钥_（如果上一步查询出来公钥，则忽略这一步）_

`ssh-keygen ``-t`` ed25519 ``-C`` ``"你的邮箱地址"`

按 Enter 键默认保存路径（`~/.ssh/id_ed25519`），无需设置密码（直接按 Enter）。

生成后，公钥文件为 `~/.ssh/id_ed25519.pub`。

3. 查看公钥内容

cat ~/.ssh/id_ed25519.pub

终端会显示类似 `ssh-ed25519 AAAA... 你的邮箱地址` 的字符串，全选复制即可。

![](static/ZfWXbcO1Pot9KrxhtpLcIz8Jncg.png)

#### **-**** ****Windows 系统**

方法一：使用 PowerShell（Windows 10/11 内置）

打开 PowerShell：右键开始菜单 → 选择 “Windows PowerShell” 或 “终端”。

生成公钥：

powershell

ssh-keygen -t ed25519 -C "你的邮箱地址"

按 Enter 键默认保存路径（`C:\Users\你的用户名\.ssh\id_ed25519`），无需密码。

查看公钥：

powershell

Get-Content C:\Users\你的用户名\.ssh\id_ed25519.pub

复制显示的字符串。

方法二：使用 Git Bash（适用于安装了 Git 的用户）

打开 Git Bash：右键点击 Git Bash 图标。

生成与查看公钥：操作同 Linux/macOS 系统，路径为 `C:/Users/你的用户名/.ssh/`。

### 登陆云开发机

#### **Web Shell 登陆**

![](static/IUasbJzyEol8j6xA6SfcI9Gsngf.png)

![](static/AuJWb3n0rodOVvxLJqzcKrSmnig.png)

Web Shell 界面类似 VSCode
将文件拖拽到左侧目录，上传文件

![](static/NBsnbzn2jo6xJIx7o0acNCk3nkd.png)

打开终端命令行进行操作

![](static/C8GUbjOrxo2FsrxTQS4cTGT3neg.png)

#### 终端登陆

点击方式 2 按钮，复制登陆命令

![](static/BJi6bZJyqohUBHxedkzcxma3ndp.png)

打开本机电脑的命令行

![](static/CFj3bb9whoYt13xEfs9cSN9Qnzh.png)

粘贴刚刚复制的命令，回车

![](static/M8kxbTU5woOMpFxqNIFcuiYInsg.png)

根据提示输入 yes，即可在本机命令行远程操作开发机

![](static/L67dbJhHoopEMBxVTFAcQNBsnDh.png)

### 体验示例工程

#### LimX Tron1 双点足

（1）下载代码包并解压至本地，拖拽上传至 Web Shell

![](static/MHyCbAas3oER8rxgGV2cosoinxe.png)

（2）打开命令行工具

![](static/Jj5Kbwy33oGrd9xLPGCcz2ZTnCf.png)

（3）使用命令行进入 train.py 所在的文件夹

命令：`cd limx`

![](static/Xf03badKYoNWN0xw3N9cefu5nVz.png)

（4）进入 limx_rl_v-main，安装下 setup.py

命令 1：`cd ..  ` 退回 limx 根目录

命令 2：`cd limx_rl_v-main/` 进入 limx_rl_v-main

命令 3：`pip install -e .`

![](static/OXP2b7qrtozwlOxniwPcsGq4nYf.png)

![](static/DttabaKvyoTiOGxIlTTcBYU8nTe.png)

（5）进入 limx/limx_legged_gym-main ，安装下 setup.py

命令 1：`cd limx`

命令 2：`cd limx_legged_gym-main`

命令 3：`pip install -e . -i ``https://pypi.tuna.tsinghua.edu.cn/simple`` --trusted-host ``pypi.tuna.tsinghua.edu.cn`

（6）尝试运行任务启动命令

命令 1：`cd ..  ` 退回 limx 根目录

命令 2：`nohu``p`` python limx_legged_gym-main/limx_legged_gym/scripts/train.py --task=pointfoot_rough --headless ``&`

出现迭代则说明成功启动训练 cd

**注：**

正常输入 python 开头的命令，如果关掉 webshell 连接，会导致 python 的进程会进入休眠。为了保证关掉 webshell 页面后，程序还在后台运行，需要在启动命令前后加上 nohup  正常的启动命令 &

![](static/WN3EbOypNoyHQXxzejBcpHNRn6e.png)

（7）获取.pt 模型文件

logs 文件夹中会有生成 model.pt 的模型文件，恭喜 🎉 你已经成功运行

![](static/Rb3bbDITVoK8C0xjDMYcXXsQnef.png)

（8）终止任务

快捷键： ctrl + C

---

**这时候，如果想再训练其他示例工程，则新上传代码，新开命令行端**

#### 本末科技 TITA

（1）上传代码包，进入文件目录

命令：`cd tita_rl-master/`

![](static/OmXCbMxXFo2Z8qxW1VRcRrofnqb.png)

（2）启动训练任务

命令：`python train.py --task=tita_constraint --headless`

![](static/Wu1Kb3KaPo2BDUxBltTcieXXnAb.png)

（3）查看模型文件

![](static/LgN3bGTb7ozlLrxWJV5cmRqanIg.png)

（4）结束训练任务

快捷键：Ctrl + C

`from setuptools import setup, find_packages`

`setup(`

`    name="tita_rl",`

`    version="0.1",`

`    packages=find_packages(),`

`    install_requires=[`

`        'matplotlib',     # 修复你当前的问题`

`        'numpy',          # 如果你用了 numpy`

`        'torch',          # 如果你用了 PyTorch`

`        # 添加其他你需要的库`

`    ],`

`    description="Legged robot RL environment",`

`    author="Your Name",`

`    author_email="``your.email@example.com``",`

`    url="``https://github.com/yourusername/tita_rl``",   # 可选`

`)`

## 【场景 4】云桌面

前置：新建后登录

![](static/KfL0bqv7IotcGKxdzf5csxjqnJe.png)

![](static/Et0HbCWwZouinZxBtWVcVAG1nHe.png)

## **训练环境任务测试**

**打开终端，切换到 root 用户下进行操作，已安装 IsaacGym 训练环境在 root 用户环境下**

```bash
#切换到root用户
sudo su root
```

### 1.1 验证 Nvidia 驱动是否安装成功

**执行：nvidia-smi，正常显示了 NVIDIA 驱动版本和 GPU 信息**

![](static/Qjv6b8qgZorbnYxZN5VcIHjtnKf.png)

### 1.2 验证 Isaac Gym 是否安装成功

```c
sudo su root
conda activate pointfoot_legged_gym
cd ~/limx_rl/isaacgym/python/examples
python 1080_balls_of_solitude.py
```

![](static/RiDIbDiqCowl1mx8JZDcaCD4n6g.png)

### 1.3 验证部署环境

打开一个 Bash 终端来运行控制算法。

```c
# 激活pointfoot_deploy的conda环境
sudo su root
conda activate pointfoot_deploy

# 运行控制算法
cd ~/limx_ws && python rl-deploy-with-python/main.py
```

打开一个 Bash 终端来运行 MuJoCo 仿真器。

```c
# 激活pointfoot_deploy的conda环境
sudo su root
conda activate pointfoot_deploy

# 运行仿真器
cd ~/limx_ws && python pointfoot-mujoco-sim/simulator.py
```

打开一个 Bash 终端来运行虚拟遥控器，仿真的过程中，您可以使用虚拟遥控器来控制机器人运动

- 左摇杆：控制前进/后退/左转/右转运动； 右摇杆：可控制机器人左右横向运动。

```c
sudo su root
cd ~/limx_ws && ./pointfoot-mujoco-sim/robot-joystick/robot-joystick
```

若能看到双点足机器人在 MuJoCo 中正常运动说明配置成功。

如果您的机器人处于摔倒状态，请在启动运动控制后，单击 MuJoCo 仿真器左侧菜单栏的 Reset 按钮以重置机器人。此时，您将看到机器人恢复并开始行走。

![](static/AROhbYKiMoiKnzx1Sc7ciaPQnOg.png)

## **TRON1 点足训练案例**

**打开终端，切换到 root 用户下进行操作，已安装 IsaacGym 训练环境在 root 用户环境下**

```bash
#切换到root用户
sudo su root
```

### **2.1 开始训练：**

常用训练相关参数指令说明：

1. 命令 1：无头模式运行（推荐使用）

```c
sudo su root
conda activate pointfoot_legged_gym
cd ~/limx_rl/pointfoot-legged-gym
python legged_gym/scripts/train.py --task=pointfoot_flat --headless
```

1. 命令 2：图形模式运行

```c
sudo su root
conda activate pointfoot_legged_gym
cd ~/limx_rl/pointfoot-legged-gym
python legged_gym/scripts/train.py --task=pointfoot_flat
```

![](static/ZwBXbVQWSoHfgVxbkVqctC4pnzI.png)

### **2.2 继续训练：**

- 如果您的训练终止，可以指定一个检查点的文件接着继续训练。
- 请注意将 `--load_run` 和 `--checkpoint` 的参数替换为实际中您的参数。

```c
sudo su root
conda activate pointfoot_legged_gym
cd ~/limx_rl/pointfoot-legged-gym
python legged_gym/scripts/train.py --task=pointfoot_flat --resume --headless --load_run Dec23_17-38-22_ --checkpoint 200
```

### **2.3 查看训练情况**

1、激活 conda 环境，并启动 tensorboard

```c
sudo su root
conda activate pointfoot_legged_gym
cd ~/limx_rl/pointfoot-legged-gym
tensorboard --logdir=logs/pointfoot_flat
```

2、在浏览器地址栏输入 `http://127.0.0.1:6006`，查看训练情况。

![](static/M2TUbHeTaoT0QTxPJNpcjTiKnNb.png)

### 2.4 导出训练结果

1、完成训练后查看训练结果 默认会读取最新的 `run` 和 `checkpoint`，如需选择特定的 `run` 和 `checkpoint`，请输入 `--load_run` 和 `--checkpoint` 参数。

**--load_run：**指定要加载的训练运行的标识符（例如，训练运行的名字或 ID）。这个标识符通常与训练过程相关联，用于从 `logs` 目录中找到相应的运行记录或配置。

- 获取方式：查看 `logs` 目录：进入 `logs` 目录查看子目录或文件，这些子目录或文件通常会以训练运行的标识符命名。
- 示例路径：
- ls -l ~/limx_rl/pointfoot-legged-gym/logs/pointfoot_flat
- 您会看到类似于 `Dec23_17-38-22_` 这样的目录，`Dec23_17-38-22_` 就是 `--load_run` 参数的值。

![](static/VQAlbfMG8oKyvhxkUYpcwEXKnef.png)

**--checkpoint：**指定要加载的检查点文件。检查点文件保存了模型的中间状态，可以用于恢复训练或进行推断。

- 获取方式：查看 `logs` 目录：在 `logs` 目录下的相应 `--load_run` 目录中，通常会有保存检查点的文件。这些文件通常以 `.pt` 或类似扩展名存在，文件名可能包含训练的轮次或时间戳。
- 示例路径：
- ls -l ~/limx_rl/pointfoot-legged-gym/logs/pointfoot_flat/Dec23_17-38-22_
- 您会看到类似于 `model_200.pt` 的文件，则 `200` 就是 `--checkpoint` 参数的值。

![](static/Idk1bX1uioqs2wxZ7q5cerc7ntb.png)

使用示例：

- 请注意将 `--load_run` 和 `--checkpoint` 后的参数替换为自己训练 `logs` 目录中的。

```c
sudo su
conda activate pointfoot*legged_gym
cd ~/limx_rl/pointfoot-legged-gym
python legged_gym/scripts/play.py --task=pointfoot_flat --load_run Dec23_17-38-22* --checkpoint 200
```

![](static/CP7Kbseoyo0gsCx2A5QcSNifncc.png)

导出模型的 ONNX 格式文件：

在运行完上一步的脚本后，可以到目录：
`~/limx_rl/pointfoot-legged-gym/logs/pointfoot_flat/exported/policies` 中查看导出的文件。

```c
ls -l ~/limx_rl/pointfoot-legged-gym/logs/pointfoot_flat/exported/policies
```

![](static/ORW6bnuV1oiloqxcT1gcWBZqnLc.png)

## **通过 gazebo 进行仿真调试**

- 设置机器人型号：请参考“查看/设置机器人型号”章节，查看您的机器人型号。如果尚未设置，请按照以下步骤进行设置。
  - 通过 Shell 命令 `tree -L 1 src/robot-description/pointfoot` 列出可用的机器人类型：

```
limx@limx:~$ tree -L 1 src/robot-description/pointfoot
src/robot-description/pointfoot
├── PF_P441A
├── PF_P441B
├── PF_P441C
├── PF_P441C2
├── PF_TRON1A
├── SF_TRON1A
└── WF_TRON1A
```

- 以 PF_TRON1A（请根据实际机器人类型进行替换）为例，设置机器人型号类型：

```
sudo su root
echo 'export ROBOT_TYPE=PF_TRON1A' >> ~/.bashrc && source ~/.bashrc
```

- 请进到您的工作空间，完成编译：
  ```bash
  ```

# 如您安装了 Conda，请临时禁用 Conda 环境

# 因为 Conda 会干扰 ROS 的运行环境设置

sudo su root
conda deactivate

cd ~/limx_ws
catkin_make install

```

- 运行部署仿真：启动后，您将看到如下所示的界面，包括Gazebo仿真器和Robot Steering交互窗口。
	- 在Gazebo窗口中，您可以使用快捷键`Ctrl + Shift + R`复位机器人；
	- 您还可以通过Robot Steering交互窗口设置，发布主题为/cmd_vel，控制机器人的移动。
	```bash
# 如您安装了Conda，请临时禁用 Conda 环境
# 因为 Conda 会干扰 ROS 的运行环境设置
sudo su root
conda deactivate

# 运行 robot_hw
cd ~/limx_ws
source install/setup.bash
roslaunch robot_hw pointfoot_hw_sim.launch
```

```
![](static/W8Z4bW2Mto2yLQxdWpacClU7nJe.png)
```

- 虚拟遥控器：如果您觉得使用 `Robot Steering` 控制机器人不方便，可以使用虚拟遥控器来简化操作。以下是使用虚拟遥控器的具体步骤。
  ```bash
  ```

# 下载虚拟遥控器

sudo su root
cd ~/limx_ws
./robot-joystick/robot-joystick

```
	![](static/BjM9bMHiFo1zMexaGdEc1bP8nzh.png)
	- 这时，您可以使用虚拟遥控器来控制机器人运动。左摇杆：控制前进/后退/左转/右转运动；右遥控：可控制机器人左右横向运动。









[阿里云无影云电脑登录及IsaacGym训练环境测试](https://cwjgfm21di.feishu.cn/docx/ZoDzd8iSMo9kTHxV9TRcOzVGn2b?from=from_copylink)
[无影云电脑登录及IsaacSim+IsaacLab训练环境测试](https://cwjgfm21di.feishu.cn/docx/DrZdd8jD3oZeQXxNh0Kcx02Kn3c?from=from_copylink)



云桌面与本地电脑之间，文件上传下载

![](static/SwOabwHUko88iDxV1ihcPCBNnjd.png)





# 三、常见问题🤔

1. **GradMotion主页**无法登录（建立安全连接失败）

用户VPN如果使用虚拟IP访问阿里的服务，可能被屏蔽

a. 重启VPN

b. 更换浏览器尝试

c. 将VPN的IP地址修改为真实的代理IP

1. **GradMotion产品怎么计费消耗？**
	按分钟计费，例如：用户资产中剩余30小时4090算力资源，开启算法训练任务1小时3分钟，则用户资产剩余28小时57分钟4090算力资源
	任务和开发机状态在【排队中】时不计费，状态为【运行中】时计费
	开发机只要开启，会持续消耗算力，开发机内无任务运行时建议及时关机



1. **为什么训练后看不到仿真视频回传？**
	1. 目前isaacgym不带有视频回传功能，仿真视频回传需要项目中有相关代码
	2. 如果想看视频回传，需要在代码包中添加相关代码，可以参考官方示例工程
		https://github.com/limxdynamics/pointfoot-legged-gym
		或参考https://github.com/isaac-sim/IsaacGymEnvs
		![](static/ZyV5b2CvCoaeeCxT6YPcl2nRnwc.png)
	3. 视频文件需要保存成mp4格式，保存在训练工程目录下任意位置都行
	---
	请提交表单[反馈问题](https://cwjgfm21di.feishu.cn/share/base/form/shrcnff2GeMWUWGtwoOS7rDsj2b)或[联系GradMotion团队](https://cwjgfm21di.feishu.cn/docx/KkZhdfQIwoicoDxTDS4csF4undL?from=from_copylink)



```
