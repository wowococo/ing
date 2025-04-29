# -*- coding: utf-8 -*-

import json
import sys

reload(sys)
sys.setdefaultencoding('utf-8')

def add_display_name(json_file):
    """
    给 JSON 文件中 'fields' 数组的每个对象添加 'display_name' 属性，
    使 'display_name' 属性的值等于 'name' 属性的值。
    """
    try:
        with open(json_file, 'r') as f:
            data = json.load(f)
    except IOError:
        print("无法打开文件: {}".format(json_file))
        return
    except json.JSONDecodeError:
        print("JSON 格式错误: {}".format(json_file))
        return

    if 'fields' in data and isinstance(data['fields'], list):
        for field in data['fields']:
            if isinstance(field, dict) and 'name' in field:
                field['display_name'] = field['name']

        try:
            with open('output.json', 'w') as f:
                json.dump(data, f, indent=4, ensure_ascii=False)  # 使用 indent=4 增加可读性
            print("成功更新文件: {}".format(json_file))
        except IOError:
            print("无法写入文件: {}".format(json_file))
    else:
        print("JSON 文件中没有 'fields' 数组或 'fields' 不是一个数组。")

# 示例用法
json_file = 'input.json'  # 替换为你的 JSON 文件名
add_display_name(json_file)