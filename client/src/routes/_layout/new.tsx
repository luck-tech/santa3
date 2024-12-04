import { createFileRoute } from "@tanstack/react-router";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import TagsInput from "@/components/TagsInput";

const formSchema = z.object({
  roomname: z
    .string()
    .min(1, "ルーム名は必須です")
    .max(50, "ルーム名は50文字以内で入力してください"),
  description: z
    .string()
    .min(1, "概要は必須です")
    .max(200, "概要は200文字以内で入力してください"),
  tags: z
    .array(z.string())
    .min(1, "技術・資格は必須です")
    .max(3, "技術・資格は3つまでです"),
});

export const Route = createFileRoute("/_layout/new")({
  component: RouteComponent,
});

function RouteComponent() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      roomname: "",
      description: "",
      tags: [],
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
  }

  return (
    <div className="container mx-auto px-6 py-4 min-h-[calc(100vh-64px)] flex flex-col justify-center max-w-screen-sm">
      <h2 className="text-2xl font-bold pb-14">ルームを作成</h2>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <FormField
            control={form.control}
            name="roomname"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-base">ルーム名</FormLabel>
                <FormControl>
                  <Input placeholder="shadcn" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-base">概要</FormLabel>
                <FormControl>
                  <Textarea placeholder="ルームについての説明" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="tags"
            render={({ field }) => (
              <FormItem>
                <FormLabel className="text-base">
                  技術・資格（3つまで）
                </FormLabel>
                <FormControl>
                  <TagsInput
                    selectedTags={field.value}
                    setSelectedTags={field.onChange}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="flex justify-end">
            <Button type="submit">作成</Button>
          </div>
        </form>
      </Form>
    </div>
  );
}