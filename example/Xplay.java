package xplay;

public class Xplay {

	public static void main(String[] args) {

		// xplayctl
		XplayCtl xc = new XplayCtl("192.168.1.16", 8700);

		Rect rect = new Rect(0, 0, 1920, 1080); // 素材显示的位置与尺寸
		
		// xc.playVideo(System.currentTimeMillis(), "PLAY_VIDEO_XXXX", 10, "/root/1.mp4", rect, "landscape", 0);
		// sleep 5s 素材在同层切换(视频切视频、视频切图片、图片切图片、图片切视频)时不需要停止，直接发布下一个素材的播放指令，xplay 会自动替换
		// xc.playVideo(System.currentTimeMillis(), "PLAY_VIDEO_XXXX", 10, "/root/2.mp4", rect, "landscape", 0);

		// xc.playImage(System.currentTimeMillis(), "PLAY_IMAGE_XXXX", 10, "/root/11.jpg", rect, "landscape", 0);

		Rect datetime = new Rect(0, 0, 500, 50);
		// xc.playDateTime(System.currentTimeMillis(), "PLAY_DATETIME_XXXX", 3, datetime, "landscape", 0, "", 30, "rgba(0, 128, 0, 100%)", "rgba(0, 0, 0, 50%)", "normal", "center");
		
		Rect qrcode = new Rect(0, 0, 100, 100);
		// xc.playQRCode(System.currentTimeMillis(), "PLAY_QRCODE_XXX", 5, "欢迎使用", qrcode, "landscape", 0);
		
		xc.playToast(System.currentTimeMillis(), "PLAY_TOAST_XXX", 2, "安装成功", "landscape", 0, "success", 15);
		
		Rect scroll = new Rect(0, 0, -1, 50); // 滚动字幕只能支持 横向全屏可自定义高 或者 竖向全屏可自定义宽 ...
		// xc.playScroll(System.currentTimeMillis(), "PLAY_SCROLL_XXX", 8, "1234567890", scroll, "landscape", 0, "", 30, "rgba(0, 128, 0, 100%)", "rgba(0, 0, 0, 50%)", "normal", "horizontal", 1);

		Rect text = new Rect(0, 0, 500, 50);
		// xc.playText(System.currentTimeMillis(), "PLAY_TEXT_XXX", 9, "55555555555", text, "landscape", 0, "", 30, "rgba(0, 128, 0, 100%)", "rgba(0, 0, 0, 50%)", "normal", "center");
		
		// xc.stopIndex(true, null);
		// xc.stopIndex(false, new ArrayList<String>(Arrays.asList("10")));
		// xc.moveIndex(10, 100, 100);
		
	}

}
