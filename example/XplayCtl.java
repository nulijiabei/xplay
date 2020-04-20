package xplay;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import java.net.Socket;
import java.net.UnknownHostException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonElement;
import com.google.gson.JsonParser;

class Rect {

	Rect(int _x, int _y, int _width, int _height) {
		x = _x;
		y = _y;
		width = _width;
		height = _height;
	}

	int x;
	int y;
	int width;
	int height;
}

class Dep {
	String path; // 素材路径
	String type; // 素材类型(video or pic)
	int duration; // 素材持续时间(非视频)
}

public class XplayCtl {

	XplayCtl(String _host, int _port) {
		host = _host;
		port = _port;
	}

	private String host;
	private int port;
	private Socket socket = null;

	private boolean connXplay(String _data) {
		try {
			if (socket == null || !socket.isConnected())
				socket = new Socket(host, port);
			OutputStreamWriter os = new OutputStreamWriter(socket.getOutputStream(), "UTF-8");
			os.write(_data + "\n#End\n");
			os.flush();
			BufferedReader br = new BufferedReader(new InputStreamReader(socket.getInputStream(), "UTF-8"));
			String data = null;
			while ((data = br.readLine()) != null) {
				if (data.contains("#End"))
					break;
				System.out.println(data);
			}
		} catch (UnknownHostException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			return false;
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			return false;
		}
		return true;
	}

	private boolean playXplay(Map<String, Object> _data) {
		Gson gson = new GsonBuilder().setPrettyPrinting().create();
		String json = gson.toJson(_data);
		JsonParser jp = new JsonParser();
		JsonElement je = jp.parse(json);
		String prettyJsonString = gson.toJson(je);
		System.out.print(prettyJsonString);
		return this.connXplay(prettyJsonString);
	}

	/*
	 * 序列(视频与图片) 
	 * _content 一组素材，每一个素材为一个 Dep 包含素材路径、类型(libName)、持续时间(非视频)
	 */
	public boolean playSequence(long _start, String _id, int _index, ArrayList<Dep> _content, Rect _rect, String _mode,
			int _rotate) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "sequence");
		data.put("start", _start);
		data.put("libName", "video");
		params.put("zIndex", _index);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		data.put("params", params);
		data.put("deps", _content);
		return this.playXplay(data);
	}

	/*
	 * 视频或流媒体等 ... 
	 * _start 该素材开始播放的时间精确到毫秒 (视频预加载问题:
	 * 提前500ms-1000ms发送指令xplay可以预加载视频)(也就是_start时间是12:00:00则在11:59:59把指令发送xplay) 
	 * _id 该素材唯一标识符(调试使用) 
	 * _index 该素材所使用的层(0-999)(层数越小越靠前，一般1-9层作为保留层，显示一些提示或者LOGO，通知等信息使用 ...) 
	 * _content 素材存储路径
	 * _rect 素材显示位置与尺寸(素材可以任意缩放拉伸，任意坐标显示) 
	 * _mode 屏幕模式(横屏：landscape、竖屏：portrait)
	 * _rotate 旋转角度(横屏：0，180、竖屏：90，270)
	 */
	public boolean playVideo(long _start, String _id, int _index, String _content, Rect _rect, String _mode,
			int _rotate) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "video");
		params.put("zIndex", _index);
		params.put("path", _content);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 图片，参数使用与视频相同
	 */
	public boolean playImage(long _start, String _id, int _index, String _content, Rect _rect, String _mode,
			int _rotate) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "pic");
		params.put("zIndex", _index);
		params.put("path", _content);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 二维码 _content 内容为二维码中的内容
	 */
	public boolean playQRCode(long _start, String _id, int _index, String _content, Rect _rect, String _mode,
			int _rotate) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "qrcode");
		params.put("zIndex", _index);
		params.put("content", _content);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * GIF，参数使用与视频相同
	 */
	public boolean playGIF(long _start, String _id, int _index, String _content, Rect _rect, String _mode,
			int _rotate) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "gif");
		params.put("zIndex", _index);
		params.put("path", _content);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 信息提示框 _content 提示框内容 
	 * _toast_type 提示框类型(notice、success、warning、error)
	 * _duration 提示框显示时间(如果为0则永久显示)
	 */
	public boolean playToast(long _start, String _id, int _index, String _content, String _mode, int _rotate,
			String _toast_type, int _duration) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "toast");
		params.put("zIndex", _index);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		params.put("content", _content);
		params.put("toast_type", _toast_type);
		params.put("duration", _duration);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 摄像头
	 *  _camera_width 与 _camera_height 为摄像头采集画面分辨率，建议(1280x720)
	 */
	public boolean playCamera(long _start, String _id, int _index, String _content, Rect _rect, String _mode,
			int _rotate, int _camera_width, int _camera_height) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "camera");
		params.put("zIndex", _index);
		params.put("device", _content);
		params.put("camera_width", _camera_width);
		params.put("camera_height", _camera_height);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 文本 
	 * _color "rgba(0, 128, 0, 100%)" 指定文本颜色RGB及透明度A 
	 * _bgcolor rgba(0, 0, 0, 0%) 指定背景颜色RGB及透明度 
	 * _style 文本样式，支持 normal、bold、italic、underline、strikethrough
	 * _align 对齐方式，支持 center、right、left _ptsize 字体大小 _ttf 指定字体文件，空则使用默认字体
	 */
	public boolean playText(long _start, String _id, int _index, String _content, Rect _rect, String _mode, int _rotate,
			String _ttf, int _ptsize, String _color, String _bgcolor, String _style, String _align) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "text");
		params.put("zIndex", _index);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		params.put("content", _content);
		params.put("color", _color);
		params.put("bgcolor", _bgcolor);
		params.put("font_size", _ptsize);
		params.put("align", _align);
		params.put("style", _style);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 滚动字幕 滚动字幕只能支持 横向全屏可自定义高 或者 竖向全屏可自定义宽 ... 
	 * _orientation 移动方向 horizontal、vertical 
	 * _speed 移动速度 默认为 1, 也就是每帧移动一个像素 ，如果 FPS30 也就是 _speed = 1 每秒 30个像素
	 */
	public boolean playScroll(long _start, String _id, int _index, String _content, Rect _rect, String _mode,
			int _rotate, String _ttf, int _ptsize, String _color, String _bgcolor, String _style, String _orientation,
			int _speed) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "scroll");
		params.put("zIndex", _index);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		params.put("content", _content);
		params.put("color", _color);
		params.put("bgcolor", _bgcolor);
		params.put("font_size", _ptsize);
		params.put("style", _style);
		params.put("orientation", _orientation);
		params.put("speed", _speed);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 日期时间
	 */
	public boolean playDateTime(long _start, String _id, int _index, Rect _rect, String _mode, int _rotate, String _ttf,
			int _ptsize, String _color, String _bgcolor, String _style, String _align) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("id", _id);
		data.put("type", "play");
		data.put("start", _start);
		data.put("libName", "datetime");
		params.put("zIndex", _index);
		params.put("left", _rect.x);
		params.put("top", _rect.y);
		params.put("width", _rect.width);
		params.put("height", _rect.height);
		params.put("screen_mode", _mode);
		params.put("screen_rotate", _rotate);
		params.put("color", _color);
		params.put("bgcolor", _bgcolor);
		params.put("font_size", _ptsize);
		params.put("align", _align);
		params.put("style", _style);
		params.put("font_ttf", _ttf);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 移动当前层坐标 ...
	 */
	public boolean moveIndex(int _index, int _x, int _y) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("type", "move");
		params.put("zIndex", _index);
		params.put("left", _x);
		params.put("top", _y);
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 停止全部层或者指定层 ...
	 */
	public boolean stopIndex(boolean _all, ArrayList<String> _indexs) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("type", "stop");
		if (_all) {
			params.put("all", true);
		} else {
			params.put("ids", _indexs);
		}
		data.put("params", params);
		return this.playXplay(data);
	}

	/*
	 * 截屏 
	 * _path 截屏文件所保存到的路径(例如:/dev/shm/snap.jpg)
	 */
	public boolean snapXplay(String _path) {
		Map<String, Object> data = new HashMap<String, Object>();
		Map<String, Object> params = new HashMap<String, Object>();
		data.put("type", "snap");
		params.put("path", _path);
		data.put("params", params);
		return this.playXplay(data);
	}

}
