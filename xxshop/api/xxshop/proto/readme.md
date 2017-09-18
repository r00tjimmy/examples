
###说明

从 api proto 文件的这两行

```

import "github.com/micro/examples/xxshop/srv/profile/proto/profile.proto";
import "github.com/micro/examples/xxshop/srv/rate/proto/rate.proto";


```


看得出， 是需要先写了 srv 的代码之后，再去定义 api 的proto文件，然后写 api 的代码的，这
才是正确的流程。


