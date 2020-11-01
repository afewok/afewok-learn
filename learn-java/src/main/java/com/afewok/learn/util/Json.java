package com.afewok.learn.util;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.*;
import lombok.extern.slf4j.Slf4j;
import org.springframework.http.converter.json.Jackson2ObjectMapperBuilder;

import java.text.SimpleDateFormat;
import java.util.List;
import java.util.Locale;
import java.util.TimeZone;

/**
 * 采用Jackson处理JSON
 */
@Slf4j
public class Json {
    private static ObjectMapper objectMapper;

    static {
        objectMapper = Jackson2ObjectMapperBuilder.json().build();
        objectMapper.setPropertyNamingStrategy(PropertyNamingStrategy.LOWER_CAMEL_CASE);//设置属性命名策略,驼峰命名
        objectMapper.setDateFormat(new SimpleDateFormat("yyyy-MM-dd HH:mm:ss"));//序列化时间默认格式
        objectMapper.setLocale(Locale.SIMPLIFIED_CHINESE);//当地时区
        objectMapper.setTimeZone(TimeZone.getTimeZone("GMT+8"));//设置为东八区
        objectMapper.setDefaultPropertyInclusion(JsonInclude.Include.ALWAYS);//空值是否序列化


//        objectMapper.configure(JsonGenerator.Feature.AUTO_CLOSE_TARGET, true);//自动关闭输出流
//        objectMapper.configure(JsonGenerator.Feature.AUTO_CLOSE_JSON_CONTENT, true);
//        objectMapper.configure(JsonGenerator.Feature.FLUSH_PASSED_TO_STREAM, true);
//        objectMapper.configure(JsonGenerator.Feature.QUOTE_FIELD_NAMES, true);
//        objectMapper.configure(JsonGenerator.Feature.QUOTE_NON_NUMERIC_NUMBERS, true);
//        objectMapper.configure(JsonGenerator.Feature.WRITE_NUMBERS_AS_STRINGS, false);
//        objectMapper.configure(JsonGenerator.Feature.WRITE_BIGDECIMAL_AS_PLAIN, false);
//        objectMapper.configure(JsonGenerator.Feature.ESCAPE_NON_ASCII, false);
//        objectMapper.configure(JsonGenerator.Feature.STRICT_DUPLICATE_DETECTION, false);
        objectMapper.configure(JsonGenerator.Feature.IGNORE_UNKNOWN, false);//在JSON中允许未引用的字段名


//        objectMapper.configure(MapperFeature.USE_ANNOTATIONS, true);//是否开启注解
//        objectMapper.configure(MapperFeature.USE_GETTERS_AS_SETTERS, true);//使用getter方法来作为setter方法
//        objectMapper.configure(MapperFeature.PROPAGATE_TRANSIENT_MARKER, false);
//        objectMapper.configure(MapperFeature.AUTO_DETECT_CREATORS, true);//是否使用creator方法来根据公共构造函数以及名字为“valueOf”的静态单参数方法自动检测
//        objectMapper.configure(MapperFeature.AUTO_DETECT_FIELDS, true);//是否非静态字段被当做属性
//        objectMapper.configure(MapperFeature.AUTO_DETECT_GETTERS, true);//是否使用“getter”方法来根据标准bean命名转换方式来自动检测
//        objectMapper.configure(MapperFeature.AUTO_DETECT_IS_GETTERS, true);//获取getter方法，前缀为is
//        objectMapper.configure(MapperFeature.AUTO_DETECT_SETTERS, true);//是否使用“setter”方法来根据标准bean命名转换方式来自动检测
        objectMapper.configure(MapperFeature.REQUIRE_SETTERS_FOR_GETTERS, true);//如果这个配置为false时，只有存在对应的构造器、setter或者field时，才调用getter。为true则忽略getter
//        objectMapper.configure(MapperFeature.ALLOW_FINAL_FIELDS_AS_MUTATORS, true);
//        objectMapper.configure(MapperFeature.INFER_PROPERTY_MUTATORS, true);
//        objectMapper.configure(MapperFeature.INFER_CREATOR_FROM_CONSTRUCTOR_PROPERTIES, true);
//        objectMapper.configure(MapperFeature.CAN_OVERRIDE_ACCESS_MODIFIERS, true);
//        objectMapper.configure(MapperFeature.OVERRIDE_PUBLIC_ACCESS_MODIFIERS, true);
//        objectMapper.configure(MapperFeature.USE_STATIC_TYPING, false);//是否使用运行时动态类型
//        objectMapper.configure(MapperFeature.USE_BASE_TYPE_AS_DEFAULT_IMPL, false);
        objectMapper.configure(MapperFeature.DEFAULT_VIEW_INCLUSION, false);//拥有view注解的属性是否在JSON序列化视图中
//        objectMapper.configure(MapperFeature.SORT_PROPERTIES_ALPHABETICALLY, false);//是否对属性使用排序，默认排序按照字母顺序
//        objectMapper.configure(MapperFeature.ACCEPT_CASE_INSENSITIVE_PROPERTIES, false);//反序列化是否忽略大小写
//        objectMapper.configure(MapperFeature.ACCEPT_CASE_INSENSITIVE_ENUMS, false);
//        objectMapper.configure(MapperFeature.USE_WRAPPER_NAME_AS_PROPERTY_NAME, false);
//        objectMapper.configure(MapperFeature.USE_STD_BEAN_NAMING, false);
//        objectMapper.configure(MapperFeature.ALLOW_EXPLICIT_PROPERTY_RENAMING, false);
//        objectMapper.configure(MapperFeature.ALLOW_COERCION_OF_SCALARS, true);
//        objectMapper.configure(MapperFeature.IGNORE_DUPLICATE_MODULE_REGISTRATIONS, true);
//        objectMapper.configure(MapperFeature.IGNORE_MERGE_FOR_UNMERGEABLE, true);


//        objectMapper.configure(DeserializationFeature.USE_BIG_DECIMAL_FOR_FLOATS, false);//浮点小数转化为BigDecimal型
//        objectMapper.configure(DeserializationFeature.USE_BIG_INTEGER_FOR_INTS, false);//整型转化为BigInteger
//        objectMapper.configure(DeserializationFeature.USE_LONG_FOR_INTS, false);//整型转化为Long
//        objectMapper.configure(DeserializationFeature.USE_JAVA_ARRAY_FOR_JSON_ARRAY, false);//JSONArray转换为转换为数组
        objectMapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);// 反序列化时,遇到未知属性时是否引起结果失败
//        objectMapper.configure(DeserializationFeature.FAIL_ON_NULL_FOR_PRIMITIVES, false);//基础数据类型（int等）为null是否报错,默认为不报错,并复制为初始值
//        objectMapper.configure(DeserializationFeature.FAIL_ON_NUMBERS_FOR_ENUMS, true);//是否根据int代表Enum的order()來反序列化Enum
//        objectMapper.configure(DeserializationFeature.FAIL_ON_INVALID_SUBTYPE, true);//反序列化时，遇到类名错误或者map中id找不到时是否报异常
//        objectMapper.configure(DeserializationFeature.FAIL_ON_READING_DUP_TREE_KEY, false);//反序列化时，遇到json数据存在两个相同的key时是否报异常
//        objectMapper.configure(DeserializationFeature.FAIL_ON_IGNORED_PROPERTIES, false);//反序列化时，遇到json属性字段为可忽略的是否报异常
//        objectMapper.configure(DeserializationFeature.FAIL_ON_UNRESOLVED_OBJECT_IDS, true);
//        objectMapper.configure(DeserializationFeature.FAIL_ON_MISSING_CREATOR_PROPERTIES, false);//Json中缺省构造函数中的参数是否进行报错
//        objectMapper.configure(DeserializationFeature.FAIL_ON_NULL_CREATOR_PROPERTIES, false);//Json中构造函数的参数为null是否进行报错
//        objectMapper.configure(DeserializationFeature.FAIL_ON_MISSING_EXTERNAL_TYPE_ID_PROPERTY, true);
//        objectMapper.configure(DeserializationFeature.FAIL_ON_TRAILING_TOKENS, false);
//        objectMapper.configure(DeserializationFeature.WRAP_EXCEPTIONS, true);//封装所有异常
//        objectMapper.configure(DeserializationFeature.ACCEPT_SINGLE_VALUE_AS_ARRAY, false);//反序列化时，是否接受单个value转为List
//        objectMapper.configure(DeserializationFeature.UNWRAP_SINGLE_VALUE_ARRAYS, false);//Json数组是否映射为实体单值
//        objectMapper.configure(DeserializationFeature.UNWRAP_ROOT_VALUE, false);//对于设置了@JsonRootName实体转换的json，可通过该属性进行实体转换
//        objectMapper.configure(DeserializationFeature.ACCEPT_EMPTY_STRING_AS_NULL_OBJECT, false);//反序列化时，是否将空字符传转化为null
//        objectMapper.configure(DeserializationFeature.ACCEPT_EMPTY_ARRAY_AS_NULL_OBJECT, false);//反序列化时，是否将空数组转换为null
//        objectMapper.configure(DeserializationFeature.ACCEPT_FLOAT_AS_INT, true);//反序列化时，是否支持Float类型转为Int
//        objectMapper.configure(DeserializationFeature.READ_ENUMS_USING_TO_STRING, false);//是否通过toString的值获取枚举对象，默认使用name()
//        objectMapper.configure(DeserializationFeature.READ_UNKNOWN_ENUM_VALUES_AS_NULL, false);//是否支持读取未知枚举，不支持则报错，支持则枚举为null
//        objectMapper.configure(DeserializationFeature.READ_UNKNOWN_ENUM_VALUES_USING_DEFAULT_VALUE, false);//读取未知枚举是否使用默认值，不使用则报错
//        objectMapper.configure(DeserializationFeature.READ_DATE_TIMESTAMPS_AS_NANOSECONDS, true);//反序列化时，是否使用long解析时间
//        objectMapper.configure(DeserializationFeature.ADJUST_DATES_TO_CONTEXT_TIME_ZONE, true);
//        objectMapper.configure(DeserializationFeature.EAGER_DESERIALIZER_FETCH, true);


//        objectMapper.configure(SerializationFeature.WRAP_ROOT_VALUE, false);//输出json格式的时候是否携带根元素
//        objectMapper.configure(SerializationFeature.INDENT_OUTPUT, false);//输出的时是否打印出漂亮的json格式
        objectMapper.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);//空对象的时候，序列化是否异常
//        objectMapper.configure(SerializationFeature.FAIL_ON_SELF_REFERENCES, true);//自我引用则失败
//        objectMapper.configure(SerializationFeature.WRAP_EXCEPTIONS, true);//封装所有异常
//        objectMapper.configure(SerializationFeature.FAIL_ON_UNWRAPPED_TYPE_IDENTIFIERS, true);
//        objectMapper.configure(SerializationFeature.CLOSE_CLOSEABLE, false);//序列化root级对象的实现closeable接口的close方法是否在序列化后被调用
//        objectMapper.configure(SerializationFeature.FLUSH_AFTER_WRITE_VALUE, true);//是否在writeValue()方法之后就调用JsonGenerator.flush()方法
        objectMapper.configure(SerializationFeature.WRITE_DATES_AS_TIMESTAMPS, false);//返回的java.util.date转换成timestamp
//        objectMapper.configure(SerializationFeature.WRITE_DATE_KEYS_AS_TIMESTAMPS, false);//将Map中得key为Date的值，也序列化为timestamps形式
//        objectMapper.configure(SerializationFeature.WRITE_DATES_WITH_ZONE_ID, false);
//        objectMapper.configure(SerializationFeature.WRITE_DURATIONS_AS_TIMESTAMPS, true);
//        objectMapper.configure(SerializationFeature.WRITE_CHAR_ARRAYS_AS_JSON_ARRAYS, false);//序列化char[]时以json数组输出
//        objectMapper.configure(SerializationFeature.WRITE_ENUMS_USING_TO_STRING, false);//序列化枚举是以toString()来输出，默认false，即默认以name()来输出
//        objectMapper.configure(SerializationFeature.WRITE_ENUMS_USING_INDEX, false);//枚举值是否序列化为数字（true）或者文本值（false）.如果是值的话，则使用Enum.ordinal()，高于WRITE_ENUMS_USING_TO_STRING
//        objectMapper.configure(SerializationFeature.WRITE_NULL_MAP_VALUES, false);//禁止把POJO中值为null的字段映射到json字符串中
//        objectMapper.configure(SerializationFeature.WRITE_EMPTY_JSON_ARRAYS, true);
//        objectMapper.configure(SerializationFeature.WRITE_SINGLE_ELEM_ARRAYS_UNWRAPPED, false);//序列化单元素数组时不以数组来输出
//        objectMapper.configure(SerializationFeature.WRITE_BIGDECIMAL_AS_PLAIN, false);
//        objectMapper.configure(SerializationFeature.WRITE_DATE_TIMESTAMPS_AS_NANOSECONDS, true);//是否将基于Date的值序列化为纳秒级别
//        objectMapper.configure(SerializationFeature.ORDER_MAP_ENTRIES_BY_KEYS, false);//序列化Map时对key进行排序操作，默认false（按hashCode降序排序）
//        objectMapper.configure(SerializationFeature.EAGER_SERIALIZER_FETCH, true);
//        objectMapper.configure(SerializationFeature.USE_EQUALITY_FOR_OBJECT_ID, false);


//        objectMapper.configure(JsonParser.Feature.AUTO_CLOSE_SOURCE, true);//是否将自动关闭那些不属于parser自己的输入源
//        objectMapper.configure(JsonParser.Feature.ALLOW_COMMENTS, true);//是否允许解析使用Java/C++ 样式的注释
//        objectMapper.configure(JsonParser.Feature.ALLOW_YAML_COMMENTS, true);//是否可以解析含有以"#"开头并直到一行结束的注释样式
//        objectMapper.configure(JsonParser.Feature.ALLOW_UNQUOTED_FIELD_NAMES, true);//是否将允许使用非双引号属性名字
//        objectMapper.configure(JsonParser.Feature.ALLOW_SINGLE_QUOTES, true);//是否允许单引号来包住属性名称和字符串值
//        objectMapper.configure(JsonParser.Feature.ALLOW_UNQUOTED_CONTROL_CHARS, true);//是否允许JSON字符串包含非引号控制字符(如：ASCII<32，如包含tab或换行符)
//        objectMapper.configure(JsonParser.Feature.ALLOW_BACKSLASH_ESCAPING_ANY_CHARACTER, true);//解析反斜杠引用的所有字符
//        objectMapper.configure(JsonParser.Feature.ALLOW_NUMERIC_LEADING_ZEROS, true);//解析以"0"为开头的数字(如: 000001)，解析时则忽略0
//        objectMapper.configure(JsonParser.Feature.ALLOW_NON_NUMERIC_NUMBERS, true);//解析非数值的数值格式（如：正无穷大，负无穷大，除以0等其他类似数据格式和xml的标签等）
//        objectMapper.configure(JsonParser.Feature.ALLOW_MISSING_VALUES, true);//JSON数组中不解析漏掉的值["value1",,"value3",]（false为抛异常，true为["value1", null, "value3", null]）
//        objectMapper.configure(JsonParser.Feature.ALLOW_TRAILING_COMMA, false);//JSON数组中是否解析带逗号的数据
//        objectMapper.configure(JsonParser.Feature.STRICT_DUPLICATE_DETECTION, false);//是否检测JSON对象重复的字段名
//        objectMapper.configure(JsonParser.Feature.IGNORE_UNDEFINED, true);//底层的数据流(二进制数据持久化，如：图片，视频等)(false为抛异常,true则忽略 )
//        objectMapper.configure(JsonParser.Feature.INCLUDE_SOURCE_IN_LOCATION, true);

    }

    public static ObjectMapper getObjectMapper() {
        return objectMapper;
    }

    /**
     * 使用泛型方法，把json字符串转换为相应的JavaBean对象。
     * (1)转换为普通JavaBean：readValue(json,Student.class)
     * (2)转换为List,如List<Student>,将第二个参数传递为Student[].class.然后使用Arrays.asList();方法把得到的数组转换为特定类型的List
     *
     * @param content
     * @param valueType
     * @return
     */
    public static <T> T parseObject(String content, Class<T> valueType) {
        try {
            return objectMapper.readValue(content, valueType);
        } catch (Exception e) {
            log.error("json对象转obj异常 e=", e);
        }
        return null;
    }

    /**
     * List<MyBean> beanList = JsonBinder.parseObject(listString, new TypeReference<List<MyBean>>() {});
     *
     * @param content
     * @param valueTypeRef
     * @return
     */
    public static <T> T parseRef(String content, TypeReference<T> valueTypeRef) {
        try {
            return objectMapper.readValue(content, valueTypeRef);
        } catch (Exception e) {
            log.error("json转obj异常 e=", e);
        }
        return null;
    }

    /**
     * Str转list
     *
     * @param content
     * @param valueType
     * @param <T>
     * @return
     */
    public static <T> List<T> parseArray(String content, Class<T> valueType) {
        try {
            JavaType javaType = objectMapper.getTypeFactory().constructParametricType(List.class, valueType);
            return objectMapper.readValue(content, javaType);
        } catch (Exception e) {
            log.error("json对象转list异常 e=", e);
        }
        return null;
    }

    /**
     * 把JavaBean转换为json字符串
     *
     * @param value
     * @return
     */
    public static String toJSONString(Object value) {
        try {
            return objectMapper.writeValueAsString(value);
        } catch (Exception e) {
            log.error("Obj转json异常 e=", e);
        }
        return null;
    }

}
