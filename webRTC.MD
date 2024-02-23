## 1. 方案
1. P2P
2. Mesh: 上下行压力大.6个人以内. 节省流量. 客户端压力大
3. MCU:MultiPoint Control Unit.合成.支持更多端.server资源消耗大,需要较高配置,节省带宽

连麦:A==WebRTC==B ------> MCU合成 ------> RTMP/CDN --------->RTMP/HLS

4. SFU:Selective Forwarding Unit.只转发,不合成. 每端1上行 n-1 下行 . server压力小,出口带宽要求较高

