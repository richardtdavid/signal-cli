import * as path from "path";
import express from "express";
import bodyParser from "body-parser";
import notifier, { NotificationMetadata } from "node-notifier";

const app = express();
const port = process.env.PORT || 9000;

app.use(bodyParser.json());

app.get("/health", (req, res) =>
  res.status(200).send({
    date: new Date(),
    status: 200,
    msg: "Signal Notifier Server Healthy",
  })
);
app.post("/notify", (req, res) => {
  notify(req.body, (reply: NotificationMetadata) => res.send(reply));
});

app.listen(port, () => console.log(`Server running on port: ${port}`));

function notify({ title, msg }: Notification, cb: CallBack) {
  console.log({ title, msg });

  notifier.notify(
    {
      title: title || "Signal",
      message: msg || "Signal",
      icon: path.join(__dirname, "signal-icon.png"),
      sound: true,
      wait: true,
      reply: true,
      closeLabel: "Done!",
      timeout: 15,
    },
    (err, response, metadata) => {
      cb(<NotificationMetadata>metadata);
    }
  );
}

type CallBack = (reply: NotificationMetadata) => void;

interface Notification {
  title: string;
  msg: string;
}
