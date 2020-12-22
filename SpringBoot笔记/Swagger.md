> url: localhost:8083/api/swagger-ui.html

### 1 配置swagger类

```java
@Configuration
@EnableSwagger2
public class Swagger2 {

    /**
     * 配置基本信息
     */
    @Bean
    public Docket createRestApi() {
        return new Docket(DocumentationType.SWAGGER_2).apiInfo(apiInfo()).select()
                .apis(RequestHandlerSelectors.basePackage("com.tower.system.controller"))
                .paths(PathSelectors.any()).build();
    }

    /**
     * 配置文档信息
     */
    private  ApiInfo apiInfo() {
        return new ApiInfoBuilder()
                .title("")
                .contact(new Contact("hankai","xxx","1197617594@qq.com"))
                .description("铁塔系统后台接口文档")
                .version("1.0").build();
    }

}
```

### 2 添加描述信息

- controller:     @Api(value = "EventSiteCheckController", tags = {"站址状态-屏蔽记录"})

- 接口： @ApiOperation(value = "上站信息查询", notes = "查询接口")

- 接收实体类: 

  - ```java
    @ApiModel(value = "查询对象", description = "这里是查询字段")
    ```

  - ```java
    @ApiModelProperty(value = "地名", name = "region", example = "湖北省", required = true) //必填字段
    ```

  - ```java
  	@ApiModelProperty(hidden = true) //隐藏字段
  	```





