## 简单工厂模式

有这么一家手机代工厂，刚开始的时候规模比较小，就是只生产安卓机，
客户需要就去工厂提货就好了。

过了几年工厂规模做的大了。开始有了新的业务了，
这家工厂不仅可以生产安卓手机，还能生产苹果手机和其他系统的手机型了。

这个时候，再像以前一样直接去工厂拿肯定是不行了，万一要的是安卓，拿的是苹果，工厂老板不得亏的底裤都没了。

这个情况，客户就必须要和工厂分开了。客户需要什么，直接和工厂负责人提就行了。由工厂负责人去把需求报给工厂，工厂给发货。

工厂模式：定义一个创建实例的接口，将“使用者”和“工厂“分开，“工厂”生产什么取决于“使用者”需要什么。

场景：一般用于对于不同实例的创建条件有着明确的不同的需求时。

优点：

    1. “使用者”想要什么东西，知道名字就行，不需要关注其他的事情。
    2. “工厂”如果需要增加什么新的业务，直接在现有基础上拓展即可，不需要对原有的业务线开刀。


        
